package gb

import (
	"reflect"
	"testing"
)

func TestRegisters_combine(t *testing.T) {
	type fields struct {
		A     byte
		B     byte
		C     byte
		D     byte
		E     byte
		F     byte
		H     byte
		L     byte
		Flags byte
	}
	type args struct {
		i byte
		j byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint16
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Registers{
				A:     tt.fields.A,
				B:     tt.fields.B,
				C:     tt.fields.C,
				D:     tt.fields.D,
				E:     tt.fields.E,
				F:     tt.fields.F,
				H:     tt.fields.H,
				L:     tt.fields.L,
				Flags: tt.fields.Flags,
			}
			if got := r.combine(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Registers.combine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegisters_set(t *testing.T) {
	type fields struct {
		A     byte
		B     byte
		C     byte
		D     byte
		E     byte
		F     byte
		H     byte
		L     byte
		Flags byte
	}
	type args struct {
		reg register
		i   uint16
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Registers{
				A:     tt.fields.A,
				B:     tt.fields.B,
				C:     tt.fields.C,
				D:     tt.fields.D,
				E:     tt.fields.E,
				F:     tt.fields.F,
				H:     tt.fields.H,
				L:     tt.fields.L,
				Flags: tt.fields.Flags,
			}
			r.set(tt.args.reg, tt.args.i)
		})
	}
}

func TestRegisters_AF(t *testing.T) {
	type fields struct {
		A     byte
		B     byte
		C     byte
		D     byte
		E     byte
		F     byte
		H     byte
		L     byte
		Flags byte
	}
	tests := []struct {
		name   string
		fields fields
		want   uint16
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Registers{
				A:     tt.fields.A,
				B:     tt.fields.B,
				C:     tt.fields.C,
				D:     tt.fields.D,
				E:     tt.fields.E,
				F:     tt.fields.F,
				H:     tt.fields.H,
				L:     tt.fields.L,
				Flags: tt.fields.Flags,
			}
			if got := r.AF(); got != tt.want {
				t.Errorf("Registers.AF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegisters_SetAF(t *testing.T) {
	type fields struct {
		A     byte
		B     byte
		C     byte
		D     byte
		E     byte
		F     byte
		H     byte
		L     byte
		Flags byte
	}
	type args struct {
		i uint16
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Registers{
				A:     tt.fields.A,
				B:     tt.fields.B,
				C:     tt.fields.C,
				D:     tt.fields.D,
				E:     tt.fields.E,
				F:     tt.fields.F,
				H:     tt.fields.H,
				L:     tt.fields.L,
				Flags: tt.fields.Flags,
			}
			r.SetAF(tt.args.i)
		})
	}
}

func TestRegisters_BC(t *testing.T) {
	type fields struct {
		A     byte
		B     byte
		C     byte
		D     byte
		E     byte
		F     byte
		H     byte
		L     byte
		Flags byte
	}
	tests := []struct {
		name   string
		fields fields
		want   uint16
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Registers{
				A:     tt.fields.A,
				B:     tt.fields.B,
				C:     tt.fields.C,
				D:     tt.fields.D,
				E:     tt.fields.E,
				F:     tt.fields.F,
				H:     tt.fields.H,
				L:     tt.fields.L,
				Flags: tt.fields.Flags,
			}
			if got := r.BC(); got != tt.want {
				t.Errorf("Registers.BC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegisters_SetBC(t *testing.T) {
	type fields struct {
		A     byte
		B     byte
		C     byte
		D     byte
		E     byte
		F     byte
		H     byte
		L     byte
		Flags byte
	}
	type args struct {
		i uint16
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Registers{
				A:     tt.fields.A,
				B:     tt.fields.B,
				C:     tt.fields.C,
				D:     tt.fields.D,
				E:     tt.fields.E,
				F:     tt.fields.F,
				H:     tt.fields.H,
				L:     tt.fields.L,
				Flags: tt.fields.Flags,
			}
			r.SetBC(tt.args.i)
		})
	}
}

func TestRegisters_DE(t *testing.T) {
	type fields struct {
		A     byte
		B     byte
		C     byte
		D     byte
		E     byte
		F     byte
		H     byte
		L     byte
		Flags byte
	}
	tests := []struct {
		name   string
		fields fields
		want   uint16
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Registers{
				A:     tt.fields.A,
				B:     tt.fields.B,
				C:     tt.fields.C,
				D:     tt.fields.D,
				E:     tt.fields.E,
				F:     tt.fields.F,
				H:     tt.fields.H,
				L:     tt.fields.L,
				Flags: tt.fields.Flags,
			}
			if got := r.DE(); got != tt.want {
				t.Errorf("Registers.DE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegisters_SetDE(t *testing.T) {
	type fields struct {
		A     byte
		B     byte
		C     byte
		D     byte
		E     byte
		F     byte
		H     byte
		L     byte
		Flags byte
	}
	type args struct {
		i uint16
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Registers{
				A:     tt.fields.A,
				B:     tt.fields.B,
				C:     tt.fields.C,
				D:     tt.fields.D,
				E:     tt.fields.E,
				F:     tt.fields.F,
				H:     tt.fields.H,
				L:     tt.fields.L,
				Flags: tt.fields.Flags,
			}
			r.SetDE(tt.args.i)
		})
	}
}

func TestRegisters_HL(t *testing.T) {
	type fields struct {
		A     byte
		B     byte
		C     byte
		D     byte
		E     byte
		F     byte
		H     byte
		L     byte
		Flags byte
	}
	tests := []struct {
		name   string
		fields fields
		want   uint16
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Registers{
				A:     tt.fields.A,
				B:     tt.fields.B,
				C:     tt.fields.C,
				D:     tt.fields.D,
				E:     tt.fields.E,
				F:     tt.fields.F,
				H:     tt.fields.H,
				L:     tt.fields.L,
				Flags: tt.fields.Flags,
			}
			if got := r.HL(); got != tt.want {
				t.Errorf("Registers.HL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegisters_SetHL(t *testing.T) {
	type fields struct {
		A     byte
		B     byte
		C     byte
		D     byte
		E     byte
		F     byte
		H     byte
		L     byte
		Flags byte
	}
	type args struct {
		i uint16
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Registers{
				A:     tt.fields.A,
				B:     tt.fields.B,
				C:     tt.fields.C,
				D:     tt.fields.D,
				E:     tt.fields.E,
				F:     tt.fields.F,
				H:     tt.fields.H,
				L:     tt.fields.L,
				Flags: tt.fields.Flags,
			}
			r.SetHL(tt.args.i)
		})
	}
}

func TestRegisters_ParseFlags(t *testing.T) {
	type fields struct {
		A     byte
		B     byte
		C     byte
		D     byte
		E     byte
		F     byte
		H     byte
		L     byte
		Flags byte
	}
	tests := []struct {
		name   string
		fields fields
		want   *Flags
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Registers{
				A:     tt.fields.A,
				B:     tt.fields.B,
				C:     tt.fields.C,
				D:     tt.fields.D,
				E:     tt.fields.E,
				F:     tt.fields.F,
				H:     tt.fields.H,
				L:     tt.fields.L,
				Flags: tt.fields.Flags,
			}
			if got := r.ParseFlags(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Registers.ParseFlags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_boolToUint8(t *testing.T) {
	type args struct {
		b bool
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := boolToUint8(tt.args.b); got != tt.want {
				t.Errorf("boolToUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegisters_SetFlags(t *testing.T) {
	type fields struct {
		A     byte
		B     byte
		C     byte
		D     byte
		E     byte
		F     byte
		H     byte
		L     byte
		Flags byte
	}
	type args struct {
		f *Flags
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Registers{
				A:     tt.fields.A,
				B:     tt.fields.B,
				C:     tt.fields.C,
				D:     tt.fields.D,
				E:     tt.fields.E,
				F:     tt.fields.F,
				H:     tt.fields.H,
				L:     tt.fields.L,
				Flags: tt.fields.Flags,
			}
			r.SetFlags(tt.args.f)
		})
	}
}
