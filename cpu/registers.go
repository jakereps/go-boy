package cpu

// Register is an internal type for simple referencing of registers.
type register byte

// an enum is used to represent the register names including the 16bit values.
const (
	a register = 0x7
	b          = 0x0
	c          = 0x1
	d          = 0x2
	e          = 0x3
	h          = 0x4
	l          = 0x5
	f          = iota
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

// registers contain various combinations of values for usage in the CPU.
type registers struct {
	a byte
	b byte
	c byte
	d byte
	e byte
	h byte
	l byte
	f byte
}

// Flags contains the 4 values for the Flags register based on the byte value.
type Flags struct {
	zero      bool
	subtract  bool
	halfCarry bool
	carry     bool
}

// combine allows for the composite usage registers (2x8-bit values = 16-bit).
func (r *registers) combine(i, j byte) uint16 {
	return uint16(i)<<8 | uint16(j)
}

// set allows mapping a 16-bit value to two internal 8-bit registers.
func (r *registers) set(reg register, i uint16) {
	v1 := byte(i & 0xFF00 >> 8)
	v2 := byte(i & 0xFF)
	switch reg {
	case af:
		r.a = v1
		r.f = v2
	case bc:
		r.b = v1
		r.c = v2
	case de:
		r.d = v1
		r.e = v2
	case hl:
		r.h = v1
		r.l = v2
	}
}

// AF returns the combination of the A and F register.
func (r *registers) AF() uint16 {
	return r.combine(r.a, r.f)
}

// SetAF breaks a 16-bit value into two 8-bit values for assignment.
func (r *registers) SetAF(i uint16) {
	r.set(af, i)
}

// BC returns the combination of the A and F register.
func (r *registers) BC() uint16 {
	return r.combine(r.b, r.c)
}

// SetBC breaks a 16-bit value into two 8-bit values for assignment.
func (r *registers) SetBC(i uint16) {
	r.set(bc, i)
}

// DE returns the combination of the A and F register.
func (r *registers) DE() uint16 {
	return r.combine(r.d, r.e)
}

// SetDE breaks a 16-bit value into two 8-bit values for assignment.
func (r *registers) SetDE(i uint16) {
	r.set(de, i)
}

// HL returns the combination of the A and F register.
func (r *registers) HL() uint16 {
	return r.combine(r.h, r.l)
}

// SetHL breaks a 16-bit value into two 8-bit values for assignment.
func (r *registers) SetHL(i uint16) {
	r.set(hl, i)
}

// ParseFlags converts the byte representation of the flags to a helper struct.
func (r *registers) ParseFlags() *Flags {
	return &Flags{
		zero:      ((r.f >> zero) & 0x1) != 0,
		subtract:  ((r.f >> subtract) & 0x1) != 0,
		halfCarry: ((r.f >> halfCarry) & 0x1) != 0,
		carry:     ((r.f >> carry) & 0x1) != 0,
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
func (r *registers) SetFlags(f *Flags) {
	r.f = (boolToByte(f.zero) << zero) | (boolToByte(f.subtract) << subtract) | (boolToByte(f.halfCarry) << halfCarry) | (boolToByte(f.carry) << carry)
}
