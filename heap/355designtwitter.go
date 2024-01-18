package heap

import (
	"container/heap"
)

type tweet struct {
	tweetID   int
	timestamp int
}

type Twitter struct {
	Follower      map[int]map[int]struct{}
	PubTweet      map[int][]tweet
	lasttimestamp int
}

type TweetMaxHeap struct {
	items []heapTweet
}

type heapTweet struct {
	idx    int // index of the tweet in the user‘s tweet list
	userId int // tweet owner

	tweet tweet // tweet itself
}

func (this *TweetMaxHeap) Len() int {
	return len(this.items)
}

func (this *TweetMaxHeap) Less(i, j int) bool {
	return this.items[i].tweet.timestamp > this.items[j].tweet.timestamp
}

func (this *TweetMaxHeap) Swap(i, j int) {
	this.items[i], this.items[j] = this.items[j], this.items[i]
}

func (this *TweetMaxHeap) Push(value interface{}) {
	this.items = append(this.items, value.(heapTweet))
}

func (this *TweetMaxHeap) Pop() interface{} {
	last := this.items[len(this.items)-1]
	this.items = this.items[:len(this.items)-1]
	return last
}

func Constructor() Twitter {
	return Twitter{
		Follower:      make(map[int]map[int]struct{}),
		PubTweet:      make(map[int][]tweet),
		lasttimestamp: 0,
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	usertweets, ok := this.PubTweet[userId]
	if !ok {
		usertweets = make([]tweet, 0)
	}
	this.lasttimestamp++

	usertweets = append(usertweets, tweet{
		tweetID:   tweetId,
		timestamp: this.lasttimestamp,
	})
	this.PubTweet[userId] = usertweets
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	// 获取userID所能看到的Post
	// 先查看其关注集
	follows, ok := this.Follower[userId]
	if !ok {
		follows = make(map[int]struct{})
	}
	// 关注自己
	if _, ok = follows[userId]; !ok {
		follows[userId] = struct{}{}
	}
	// 将自己也加进自己的关注用户集里
	this.Follower[userId] = follows

	// 大顶堆
	tweetHeap := &TweetMaxHeap{items: make([]heapTweet, 0)}

	for followee := range follows {
		// 关注用户的tweets集
		tweets := this.PubTweet[followee]
		if len(tweets) == 0 {
			continue
		}
		// 将新的推文 推进大顶堆
		heap.Push(tweetHeap, heapTweet{
			tweet:  tweets[len(tweets)-1],
			idx:    len(tweets) - 1,
			userId: followee,
		})
	}

	var result []int

	for tweetHeap.Len() > 0 && len(result) < 10 {
		popped := heap.Pop(tweetHeap).(heapTweet)
		// 将tweetID 加进结果集中
		result = append(result, popped.tweet.tweetID)
		// 下一个tweet文章的索引
		nextIdx := popped.idx - 1
		//     // If there is a next tweet in the list, push it onto the max heap

		if nextIdx >= 0 {
			tweets := this.PubTweet[popped.userId]
			heap.Push(tweetHeap, heapTweet{
				tweet:  tweets[nextIdx],
				idx:    nextIdx,
				userId: popped.userId,
			})
		}
	}
	return result
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	// 获取关注者关注的用户集
	followees, ok := this.Follower[followerId]
	// 如果尚未存在于映射中,创建一个空的关注用户集合
	if !ok {
		followees = make(map[int]struct{})
		this.Follower[followerId] = followees
	}
	// 将被关注用户(followeeID)添加到关注者的关注集合中
	followees[followeeId] = struct{}{}
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	followees, ok := this.Follower[followerId]
	if !ok {
		return
	}
	delete(followees, followeeId)
}
