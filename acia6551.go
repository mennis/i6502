package i6502

const (
	aciaData = iota
	aciaStatus
	aciaCommand
	aciaControl
)

/*
ACIA 6551 Serial IO
*/
type Acia6551 struct {
	Rx chan byte // Reading (Acia Input) line
	Tx chan byte // Transmitting (Acia Output) line

	rxData      byte
	txData      byte
	commandData byte
	controlData byte

	rxFull  bool
	txEmpty bool

	rxIrqEnabled bool
	txIrqEnabled bool
}

func NewAcia6551(rx chan byte, tx chan byte) (*Acia6551, error) {
	acia := &Acia6551{Tx: tx, Rx: rx}
	acia.Reset()

	go func() {
		// Handle rx data channel
	}()

	go func() {
		// Handle tx data channel
	}()

	return acia, nil
}

func (a *Acia6551) Size() uint16 {
	// We have a only 4 addresses, Data, Status, Command and Control
	return 0x04
}

func (a *Acia6551) Reset() {
	a.rxData = 0
	a.rxFull = false

	a.txData = 0
	a.txEmpty = true

	a.rxIrqEnabled = false
	a.txIrqEnabled = false
}

/*
func (r *Rom) Read(address uint16) byte {
	return r.data[address]
}

func (r *Rom) Write(address uint16, data byte) {
	panic(fmt.Errorf("Trying to write to ROM at 0x%04X", address))
}
*/
