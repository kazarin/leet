package main

type MovingAverage struct {
	queue []int
	size  int
}

func Constructor(size int) MovingAverage {
	ma := MovingAverage{size: size}
	ma.queue = []int{}
	return ma

}

func (this *MovingAverage) Next(val int) float64 {
	this.queue = append(this.queue, val)
	if len(this.queue) > this.size {
		this.queue = this.queue[len(this.queue)-this.size:]
	}

	var sum float64 = 0.0
	for i := 0; i < len(this.queue); i++ {
		sum += float64(this.queue[i])
	}
	l := float64(max(1, len(this.queue)))
	return sum / l

}
