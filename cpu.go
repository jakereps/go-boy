package gb

import "fmt"

// register is an internal type for simple referencing of registers.
type register byte

// an enum is used to represent the register names including the 16bit values.
const (
	a register = iota + 1
	b
	c
	d
	e
	f
	h
	l
	af
	bc
	de
	hl
)

// an enum for the flag byte offset on register flags.
const (
	carry byte = iota + 4
	halfCarry
	subtract
	zero
)

// Registers contain various combinations of values for usage in the CPU.
type Registers struct {
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

// Flags contains the 4 values for the Flags register based on the byte value.
type Flags struct {
	Zero      bool
	Subtract  bool
	HalfCarry bool
	Carry     bool
}

// combine allows for the composite usage registers (2x8-bit values = 16-bit).
func (r *Registers) combine(i, j byte) uint16 {
	return uint16(i)<<8 | uint16(j)
}

// set allows mapping a 16-bit value to two internal 8-bit registers.
func (r *Registers) set(reg register, i uint16) {
	v1 := byte(i & 0xFF00 >> 8)
	v2 := byte(i & 0xFF)
	switch reg {
	case af:
		r.A = v1
		r.F = v2
	case bc:
		r.B = v1
		r.C = v2
	case de:
		r.D = v1
		r.E = v2
	case hl:
		r.H = v1
		r.L = v2
	}
}

// AF returns the combination of the A and F register.
func (r *Registers) AF() uint16 {
	return r.combine(r.A, r.F)
}

// SetAF breaks a 16-bit value into two 8-bit values for assignment.
func (r *Registers) SetAF(i uint16) {
	r.set(af, i)
}

// BC returns the combination of the A and F register.
func (r *Registers) BC() uint16 {
	return r.combine(r.B, r.C)
}

// SetBC breaks a 16-bit value into two 8-bit values for assignment.
func (r *Registers) SetBC(i uint16) {
	r.set(bc, i)
}

// DE returns the combination of the A and F register.
func (r *Registers) DE() uint16 {
	return r.combine(r.D, r.E)
}

// SetDE breaks a 16-bit value into two 8-bit values for assignment.
func (r *Registers) SetDE(i uint16) {
	r.set(de, i)
}

// HL returns the combination of the A and F register.
func (r *Registers) HL() uint16 {
	return r.combine(r.H, r.L)
}

// SetHL breaks a 16-bit value into two 8-bit values for assignment.
func (r *Registers) SetHL(i uint16) {
	r.set(hl, i)
}

// ParseFlags converts the byte representation of the flags to a helper struct.
func (r *Registers) ParseFlags() *Flags {
	fmt.Println(zero)
	fmt.Println((0x80 >> zero))
	return &Flags{
		Zero:      ((r.Flags >> zero) & 0x1) != 0,
		Subtract:  ((r.Flags >> subtract) & 0x1) != 0,
		HalfCarry: ((r.Flags >> halfCarry) & 0x1) != 0,
		Carry:     ((r.Flags >> carry) & 0x1) != 0,
	}
}

// boolToByte maps a bool to a numeric value for bitwise ops.
func boolToByte(b bool) byte {
	if b {
		return 1
	}
	return 0
}

// SetFlags assigns the register flags to match the provided Flags' values.
func (r *Registers) SetFlags(f *Flags) {
	r.Flags = (boolToByte(f.Zero) << zero) | (boolToByte(f.Subtract) << subtract) | (boolToByte(f.HalfCarry) << halfCarry) | (boolToByte(f.Carry) << carry)
}
