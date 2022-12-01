package adapter

type LeadAdapter interface {
	SetScheduler(string)
	SetGetLimit(int)
	RemoveScheduler()
}
