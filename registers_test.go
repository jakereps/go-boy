package gb

import (
	"reflect"
	"testing"
)

func TestRegisters_combine(t *testing.T) {
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
			r := &Registers{}
			if got := r.combine(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Registers.combine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegisters_set(t *testing.T) {
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
			r := &Registers{}
			r.set(tt.args.reg, tt.args.i)
			var got uint16
			switch tt.args.reg {
			case af:
				got = uint16(r.A)<<8 | uint16(r.F)
			case bc:
				got = uint16(r.B)<<8 | uint16(r.C)
			case de:
				got = uint16(r.D)<<8 | uint16(r.E)
			case hl:
				got = uint16(r.H)<<8 | uint16(r.L)
			default:
				t.Fatal("invalid reg provided in test args")
			}
			if got != tt.args.i {
				t.Errorf("Registers.set() = %v, want %v", got, tt.args.i)
			}
		})
	}
}

func TestRegisters_ParseFlags(t *testing.T) {
	type fields struct {
		F byte
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
				F: 0x80,
			},
			want: &Flags{
				Zero: true,
			},
		},
		{
			name: "subtract",
			fields: fields{
				// 0b0100 0000
				F: 0x40,
			},
			want: &Flags{
				Subtract: true,
			},
		},
		{
			name: "halfCarry",
			fields: fields{
				// 0b0010 0000
				F: 0x20,
			},
			want: &Flags{
				HalfCarry: true,
			},
		},
		{
			name: "carry",
			fields: fields{
				// 0b0001 0000
				F: 0x10,
			},
			want: &Flags{
				Carry: true,
			},
		},
		{
			name: "all",
			fields: fields{
				F: 0xFF,
			},
			want: &Flags{
				Zero:      true,
				Subtract:  true,
				HalfCarry: true,
				Carry:     true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Registers{
				F: tt.fields.F,
			}
			if got := r.ParseFlags(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Registers.ParseFlags() = %v, want %v", got, tt.want)
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

func TestRegisters_SetFlags(t *testing.T) {
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
					Zero: true,
				},
			},
		},
		{
			name: "subtract",
			args: args{
				f: &Flags{
					Subtract: true,
				},
			},
		},
		{
			name: "halfCarry",
			args: args{
				f: &Flags{
					HalfCarry: true,
				},
			},
		},
		{
			name: "carry",
			args: args{
				f: &Flags{
					Carry: true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Registers{}
			r.SetFlags(tt.args.f)
			if tt.args.f.Zero {
				if ((r.F >> zero) & 0x1) != 1 {
					t.Errorf("Registers.SetFlags().Zero = %v, want %v", ((r.F >> zero) & 0x1), 1)
				}
			}
			if tt.args.f.Subtract {
				if ((r.F >> subtract) & 0x1) != 1 {
					t.Errorf("Registers.SetFlags().Subtract = %v, want %v", ((r.F >> subtract) & 0x1), 1)
				}
			}
			if tt.args.f.HalfCarry {
				if ((r.F >> halfCarry) & 0x1) != 1 {
					t.Errorf("Registers.SetFlags().HalfCarry = %v, want %v", ((r.F >> halfCarry) & 0x1), 1)
				}
			}
			if tt.args.f.Carry {
				if ((r.F >> carry) & 0x1) != 1 {
					t.Errorf("Registers.SetFlags().Carry = %v, want %v", ((r.F >> carry) & 0x1), 1)
				}
			}
		})
	}
}
