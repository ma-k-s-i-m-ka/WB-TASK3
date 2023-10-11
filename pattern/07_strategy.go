package pattern

// Интерфейс для стратегии сортировки
type SortStrategy interface {
	Sort(data []int)
}

// Стратегия сортировки пузырьком
type BubbleSortStrategy struct {
}

func (b *BubbleSortStrategy) Sort(data []int) {
	n := len(data)
	for i := 0; i < n; i++ {
		for j := n - 1; j >= i+1; j-- {
			if data[j] < data[j-1] {
				data[j], data[j-1] = data[j-1], data[j]
			}
		}
	}
}

// Стратегия сортировки слиянием
type MergeSortStrategy struct {
}

func (m *MergeSortStrategy) Sort(data []int) {
	n := len(data)
	for i := 1; i < n; i++ {
		var j int
		var buff = data[i]
		for j = i - 1; j >= 0; j-- {
			if data[j] < buff {
				break
			}
			data[j+1] = data[j]
		}
		data[j+1] = buff
	}
}

// Контекст, использующий стратегию сортировки
type SortContext struct {
	strategy SortStrategy
}

func (c *SortContext) SetStrategy(strategy SortStrategy) {
	c.strategy = strategy
}

func (c *SortContext) Sort(data []int) {
	c.strategy.Sort(data)
}
