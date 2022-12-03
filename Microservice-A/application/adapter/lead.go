package adapter

type LeadAdapter interface {
	SetScheduler(string)
	SetGetLimit(int)
	Pause(bool)
}

type LeadRepository interface {
	GetLead() (*[]byte, error)
}
