/*
 * @Descripttion: 基于物品的协同过滤
 * @Author:downson
 * @Date: 2024-04-10 21:57:38
 * @LastEditTime: 2024-04-11 18:52:41
 */
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type User struct {
	userID int
	item   [4]Post
}

type Post struct {
	PostId   int
	PostName string
	PostTag  int // 1.星座 2.感情  3.社会
	Score    int
}

func initData() ([]*User, map[int][]Post) {
	var users []*User
	postPool := make(map[int][]Post)

	for i := 0; i < 3; i++ {
		user := &User{userID: 10000 + i}
		users = append(users, user)
	}
	for i := range users {
		if users[i].userID == 10000 {
			users[i].item[0].PostId = 10
			users[i].item[0].PostName = "水瓶座故事"
			users[i].item[0].PostTag = 1

			users[i].item[1].PostId = 11
			users[i].item[1].PostName = "清华大学教授出轨17岁女生"
			users[i].item[1].PostTag = 3

			users[i].item[2].PostId = 12
			users[i].item[2].PostName = "摩羯座和水瓶座"
			users[i].item[2].PostTag = 1

			users[i].item[3].PostId = 13
			users[i].item[3].PostName = "两个人分手如何符合"
			users[i].item[3].PostTag = 2
		}

		if users[i].userID == 10001 {
			users[i].item[0].PostId = 14
			users[i].item[0].PostName = "摩羯座故事"
			users[i].item[0].PostTag = 1

			users[i].item[1].PostId = 15
			users[i].item[1].PostName = "娜扎和张翰分手"
			users[i].item[1].PostTag = 3

			users[i].item[2].PostId = 16
			users[i].item[2].PostName = "摩羯座"
			users[i].item[2].PostTag = 1

			users[i].item[3].PostId = 17
			users[i].item[3].PostName = "两个人分手如何符合"
			users[i].item[3].PostTag = 2
		}

		if users[i].userID == 10002 {
			users[i].item[0].PostId = 18
			users[i].item[0].PostName = "星盘"
			users[i].item[0].PostTag = 1

			users[i].item[1].PostId = 20
			users[i].item[1].PostName = "xx和张翰分手"
			users[i].item[1].PostTag = 3

			users[i].item[2].PostId = 16
			users[i].item[2].PostName = "摩羯座"
			users[i].item[2].PostTag = 1

			users[i].item[3].PostId = 22
			users[i].item[3].PostName = "两个人结婚如何准备"
			users[i].item[3].PostTag = 2
		}
	}

	posts := []Post{
		{PostId: 1, PostName: "双鱼座梦境解析", PostTag: 1},
		{PostId: 2, PostName: "处女座的完美主义", PostTag: 1},
		{PostId: 3, PostName: "双鱼座的创意天赋", PostTag: 1},
		{PostId: 4, PostName: "我和男朋友的矛盾", PostTag: 2},
		{PostId: 5, PostName: "我和女朋友的矛盾", PostTag: 2},
		{PostId: 6, PostName: "我和我自己的矛盾", PostTag: 2},
		{PostId: 7, PostName: "政治", PostTag: 3},
		{PostId: 8, PostName: "军事", PostTag: 3},
		{PostId: 9, PostName: "武力", PostTag: 3},
		{PostId: 10, PostName: "水瓶座故事", PostTag: 1},
		{PostId: 11, PostName: "清华大学教授出轨17岁女生", PostTag: 3},
		{PostId: 12, PostName: "摩羯座和水瓶座", PostTag: 1},
		{PostId: 13, PostName: "两个人分手如何符合", PostTag: 2},
		{PostId: 14, PostName: "摩羯座故事", PostTag: 1},
		{PostId: 15, PostName: "娜扎和张翰分手", PostTag: 3},
		{PostId: 16, PostName: "摩羯座", PostTag: 1},
		{PostId: 17, PostName: "两个人分手如何符合", PostTag: 2},
		{PostId: 18, PostName: "星盘", PostTag: 1},
		{PostId: 20, PostName: "xx和张翰分手", PostTag: 3},
		{PostId: 22, PostName: "两个人结婚如何准备", PostTag: 2},
	}

	for _, post := range posts {
		postPool[post.PostTag] = append(postPool[post.PostTag], post)
	}
	//fmt.Println(users, postPool)
	return users, postPool
}

func user_like_items(users []*User, post Post) {
	for i := 0; i < len(users); i++ {
		if users[i].userID == 10000 {
			users[i].item[0].Score = 8
			users[i].item[1].Score = 5
			users[i].item[2].Score = 4
			users[i].item[3].Score = 1
		}
		if users[i].userID == 10001 {
			users[i].item[0].Score = 8
			users[i].item[1].Score = 1
			users[i].item[2].Score = 8
			users[i].item[3].Score = 1
		}

		if users[i].userID == 10002 {
			users[i].item[0].Score = 0
			users[i].item[1].Score = 7
			users[i].item[2].Score = 0
			users[i].item[3].Score = 8
		}
	}
}

func sim_posts(postPool map[int][]Post) map[int]map[int]float64 {
	postSimilarity := make(map[int]map[int]float64)

	for _, posts := range postPool {
		for i := 0; i < len(posts); i++ {
			for j := i + 1; j < len(posts); j++ {
				sim := cal_post_sim(posts[i], posts[j])

				if postSimilarity[posts[i].PostId] == nil {
					postSimilarity[posts[i].PostId] = make(map[int]float64)
				}
				postSimilarity[posts[i].PostId][posts[j].PostId] = sim

				if postSimilarity[posts[j].PostId] == nil {
					postSimilarity[posts[j].PostId] = make(map[int]float64)
				}
				postSimilarity[posts[j].PostId][posts[i].PostId] = sim
			}
		}
	}
	//	fmt.Println(postSimilarity)
	return postSimilarity
}

func cal_post_sim(post1, post2 Post) float64 {
	if post1.PostTag == post2.PostTag {
		return 1.0
	} else {
		return 0.0
	}
}

func rem_post_to_user(users []*User, postPool map[int][]Post, postSimilarity map[int]map[int]float64) map[int][]Post {
	rec := make(map[int][]Post, 5)

	for _, user := range users {
		viewedPosts := make(map[int]bool)
		for _, item := range user.item {
			viewedPosts[item.PostId] = true
		}
		for _, item := range user.item {
			// 检查相似度矩阵中是否存在值为1的情况
			if postSimilarity[item.PostId] != nil {
				for postId, similarity := range postSimilarity[item.PostId] {
					// 如果存在相似度为1且未被浏览的帖子，则推荐给用户
					if postId != item.PostId && similarity == 1 && !viewedPosts[postId] {
						similarPosts := postPool[postId]
						for _, post := range similarPosts {
							rec[user.userID] = append(rec[user.userID], post)
						}
					}
				}
			}
		}
	}

	return rec
}

func get_random_posts(recommendations map[int][]Post) map[int][]Post {
	rand.Seed(time.Now().UnixNano())
	randomPosts := make(map[int][]Post)

	for userID, posts := range recommendations {
		if len(posts) > 0 {
			rand.Shuffle(len(posts), func(i, j int) {
				posts[i], posts[j] = posts[j], posts[i]
			})
			numPosts := min(5, len(posts))
			randomPosts[userID] = posts[:numPosts]
		}
	}

	return randomPosts
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	users, postPool := initData()
	postSimilarity := sim_posts(postPool)
	recommendations := rem_post_to_user(users, postPool, postSimilarity)
	fmt.Println(recommendations)
	randomPosts := get_random_posts(recommendations)

	for userID, posts := range randomPosts {
		fmt.Printf("UserID: %d, Recommended Posts: %v\n", userID, posts)
	}
}
