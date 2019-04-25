package cpu

// Instruction is the representation of a Game Boy CPU instruction.
type Instruction byte

// Game Boy CPU instructions supported.
const (
	LD Instruction = iota
	LDHL
	PUSH
	POP
	ADD
	ADC
	SUB
	SBC
	AND
	OR
	XOR
	CP
	INC
	DEC
	RLCA
	RLA
	RRCA
	RRA
	RLC
	RL
	RRC
	RR
	SLA
	SRA
	SRL
	SWAP
	BIT
	SET
	RES
	JP
	JR
	CALL
	RET
	RETI
	RST
	DAA
	CPL
	NOP
	CCF
	SCF
	DI
	EI
	HALT
	STOP
)
