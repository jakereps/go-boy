package cpu

import (
	"reflect"
	"testing"
)

func Testregisters_combine(t *testing.T) {
	type args struct {
		i byte
		j byte
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{
			name: "success",
			args: args{
				i: 0x1,
				j: 0x1,
			},
			// 0b0000 0001 0000 0001
			want: 0x101,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &registers{}
			if got := r.combine(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("registers.combine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Testregisters_set(t *testing.T) {
	type args struct {
		reg register
		i   uint16
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success-af",
			args: args{
				reg: af,
				i:   257,
			},
		},
		{
			name: "success-bc",
			args: args{
				reg: bc,
				i:   128,
			},
		},
		{
			name: "success-de",
			args: args{
				reg: de,
				i:   512,
			},
		},
		{
			name: "success-hl",
			args: args{
				reg: hl,
				i:   1024,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &registers{}
			r.set(tt.args.reg, tt.args.i)
			var got uint16
			switch tt.args.reg {
			case af:
				got = uint16(r.a)<<8 | uint16(r.f)
			case bc:
				got = uint16(r.b)<<8 | uint16(r.c)
			case de:
				got = uint16(r.d)<<8 | uint16(r.e)
			case hl:
				got = uint16(r.h)<<8 | uint16(r.l)
			default:
				t.Fatal("invalid reg provided in test args")
			}
			if got != tt.args.i {
				t.Errorf("registers.set() = %v, want %v", got, tt.args.i)
			}
		})
	}
}

func Testregisters_ParseFlags(t *testing.T) {
	type fields struct {
		f byte
	}
	tests := []struct {
		name   string
		fields fields
		want   *Flags
	}{
		{
			name: "zero",
			fields: fields{
				// 0b1000 0000
				f: 0x80,
			},
			want: &Flags{
				zero: true,
			},
		},
		{
			name: "subtract",
			fields: fields{
				// 0b0100 0000
				f: 0x40,
			},
			want: &Flags{
				subtract: true,
			},
		},
		{
			name: "halfCarry",
			fields: fields{
				// 0b0010 0000
				f: 0x20,
			},
			want: &Flags{
				halfCarry: true,
			},
		},
		{
			name: "carry",
			fields: fields{
				// 0b0001 0000
				f: 0x10,
			},
			want: &Flags{
				carry: true,
			},
		},
		{
			name: "all",
			fields: fields{
				f: 0xFF,
			},
			want: &Flags{
				zero:      true,
				subtract:  true,
				halfCarry: true,
				carry:     true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &registers{
				f: tt.fields.f,
			}
			if got := r.ParseFlags(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("registers.ParseFlags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_boolToByte(t *testing.T) {
	type args struct {
		b bool
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "true",
			args: args{
				b: true,
			},
			want: 1,
		},
		{
			name: "false",
			args: args{
				b: false,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := boolToByte(tt.args.b); got != tt.want {
				t.Errorf("boolToByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Testregisters_SetFlags(t *testing.T) {
	type args struct {
		f *Flags
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "zero",
			args: args{
				f: &Flags{
					zero: true,
				},
			},
		},
		{
			name: "subtract",
			args: args{
				f: &Flags{
					subtract: true,
				},
			},
		},
		{
			name: "halfCarry",
			args: args{
				f: &Flags{
					halfCarry: true,
				},
			},
		},
		{
			name: "carry",
			args: args{
				f: &Flags{
					carry: true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &registers{}
			r.SetFlags(tt.args.f)
			if tt.args.f.zero {
				if ((r.f >> zero) & 0x1) != 1 {
					t.Errorf("registers.SetFlags().zero = %v, want %v", ((r.f >> zero) & 0x1), 1)
				}
			}
			if tt.args.f.subtract {
				if ((r.f >> subtract) & 0x1) != 1 {
					t.Errorf("registers.SetFlags().subtract = %v, want %v", ((r.f >> subtract) & 0x1), 1)
				}
			}
			if tt.args.f.halfCarry {
				if ((r.f >> halfCarry) & 0x1) != 1 {
					t.Errorf("registers.SetFlags().halfCarry = %v, want %v", ((r.f >> halfCarry) & 0x1), 1)
				}
			}
			if tt.args.f.carry {
				if ((r.f >> carry) & 0x1) != 1 {
					t.Errorf("registers.SetFlags().Carry = %v, want %v", ((r.f >> carry) & 0x1), 1)
				}
			}
		})
	}
}
