package cpu

// LR35902 simulates a Game Boy CPU through the usage of registers, program counter,
// stack pointer, and more. More or less this will be the full "logic" of a real CPU.
type LR35902 struct {
	registers *registers
	pc        uint16
	sp        uint16
}
