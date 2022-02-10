package action

const AskIntACT actionType = "ask_int"

type AskIntAction struct {
	actionComm

	from     uint32
	attacker uint32
}

type AskIntActionJSON struct {
	actionJSONComm
	From     uint32 `json:"from"`
	Attacker uint32 `json:"attacker"`
}
