package browserhistory

// https://leetcode.com/problems/design-browser-history/
type BrowserHistory struct {
	history     []string
	currentStep int
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{history: []string{homepage}, currentStep: 1}
}

func (this *BrowserHistory) Visit(url string) {
	this.history = this.history[:this.currentStep]
	this.history = append(this.history, url)
	this.currentStep++
}

func (this *BrowserHistory) Back(steps int) string {
	destIndex := this.currentStep - steps
	if destIndex <= 0 {
		destIndex = 1
	}
	this.currentStep = destIndex
	return this.history[destIndex-1]
}

func (this *BrowserHistory) Forward(steps int) string {
	destIndex := this.currentStep + steps
	maxIndex := len(this.history)
	if destIndex >= maxIndex {
		destIndex = maxIndex
	}
	this.currentStep = destIndex
	return this.history[destIndex-1]
}
