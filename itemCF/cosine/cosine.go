/*
 * @Descripttion:
 * @Author:
 * @Date: 2024-04-11 19:00:16
 * @LastEditTime: 2024-04-11 19:30:47
 */
package cosine

import "math"

type User struct {
	ID    int
	Likes []int
}

type Item struct {
	ID    int
	Likes []int
}

type UserLikes map[int][]int

func initData() ([]User, []Item, UserLikes) {
	// 示例用户数据
	users := []User{
		{ID: 1, Likes: []int{101, 102, 103}},
		{ID: 2, Likes: []int{102, 104}},
		{ID: 3, Likes: []int{101, 103, 104}},
	}

	// 示例物品数据
	items := []Item{
		{ID: 101, Likes: []int{1, 3}},
		{ID: 102, Likes: []int{1, 2}},
		{ID: 103, Likes: []int{1, 3}},
		{ID: 104, Likes: []int{2, 3}},
	}

	// 构建用户喜好的物品ID映射
	userLikes := make(UserLikes)
	for _, user := range users {
		for _, itemID := range user.Likes {
			userLikes[user.ID] = append(userLikes[user.ID], itemID)
		}
	}

	return users, items, userLikes
}

func calculateItemSimilarity(items []Item, userLikes UserLikes) map[int]map[int]float64 {
	itemSimilarity := make(map[int]map[int]float64)

	for _, item1 := range items {
		itemSimilarity[item1.ID] = make(map[int]float64)
		for _, item2 := range items {
			if item1.ID == item2.ID {
				continue
			}
			// 计算交集
			commonUsers := intersection(item1.Likes, item2.Likes)
			// 计算这两个物品在共同喜欢的用户上的分数的平方和。这些分数是用户对物品的喜好程度的评分。
			sum1 := calculateSumOfSquares(item1.ID, commonUsers, userLikes)
			sum2 := calculateSumOfSquares(item2.ID, commonUsers, userLikes)
			//	计算
			sumProduct := calculateSumProduct(item1.ID, item2.ID, commonUsers, userLikes)
			similarity := sumProduct / (math.Sqrt(sum1) * math.Sqrt(sum2))
			//itemID1 和 itemID2 是要计算乘积总和的两个物品的标识符，commonUsers 是这两个物品共同喜欢的用户的集合，
			// userLikes 是一个映射，将用户ID映射到一个映射，该映射将物品ID映射到用户对该物品的评分。
			itemSimilarity[item1.ID][item2.ID] = similarity
		}
	}

	return itemSimilarity
}

// 计算交集
func intersection(a, b []int) []int {
	set := make(map[int]bool)
	var result []int
	for _, val := range a {
		set[val] = true
	}
	for _, val := range b {
		if set[val] {
			result = append(result, val)
		}
	}
	return result
}

func calculateSumOfSquares(itemID int, commonUsers []int, userLikes UserLikes) float64 {
	sum := 0.0
	for _, userID := range commonUsers {
		sum += math.Pow(float64(len(userLikes[userID])), 2)
	}
	return sum
}

func calculateSumProduct(itemID1, itemID2 int, commonUsers []int, userLikes UserLikes) float64 {
	sum := 0.0
	for _, userID := range commonUsers {
		sum += float64(len(userLikes[userID]))
	}
	return sum
}

func main() {

}
