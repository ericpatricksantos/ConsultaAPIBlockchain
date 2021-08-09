package Model

type Database struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SSLmode  string
}

type Block struct {
	Hash          string        `json:"hash"`
	Ver           int           `json:"ver"`
	Prev_block    string        `json:"prev_block"`
	Mrkl_root     string        `json:"mrkl_root"`
	Time          int           `json:"time"`
	Bits          int           `json:"bits"`
	Next_Block    []string      `json:"next_block"`
	Nonce         int           `json:"nonce"`
	N_tx          int           `json:"n_tx"`
	Block_index   int           `json:"block_index"`
	Main_chain    bool          `json:"main_chain"`
	Height        int           `json:"height"`
	Received_time int           `json:"received_time"`
	Relayed_by    string        `json:"relayed_by"`
	Tx            []Transaction `json:"tx"`
}

type Transaction struct {
	Hash         string        `json:"hash"`
	Ver          int           `json:"ver"`
	Vin_sz       int           `json:"vin_sz"`
	Vout_sz      int           `json:"vout_sz"`
	Lock_time    string        `json:"lock_time"`
	Size         string        `json:"size"`
	Relayed_by   string        `json:"relayed_by"`
	Block_height int           `json:"block_height"`
	Tx_index     string        `json:"tx_index"`
	Inputs       []ListaInputs `json:"inputs"`
	Out          []ListaOut    `json:"out"`
}

type ListaInputs struct {
	Prev_out PrevOut `json:"prev_out"`
	Script   string  `json:"script"`
}

type ListaOut struct {
	Value  string `json:"value"`
	Hash   string `json:"hash"`
	Script string `json:"script"`
}

type PrevOut struct {
	Hash     string `json:"hash"`
	Value    string `json:"value"`
	Tx_index string `json:"tx_index"`
	N        int    `json:"n"`
}
