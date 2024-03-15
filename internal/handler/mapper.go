package handler

import (
	"github.com/abhithube/at-feeds/internal/api"
	"github.com/abhithube/at-feeds/internal/database"
)

func mapToAPICollection(collection database.Collection) api.Collection {
	result := api.Collection{
		Id:    int(collection.ID),
		Title: collection.Title,
	}

	return result
}

func mapToAPIFeed(feed database.Feed, totalEntryCount, unreadEntryCount *int64) api.Feed {
	result := api.Feed{
		Id:    int(feed.ID),
		Link:  feed.Link,
		Title: feed.Title,
	}
	if feed.Url.Valid {
		result.Url = &feed.Url.String
	}
	if totalEntryCount != nil {
		entryCount := int(*totalEntryCount)
		result.TotalEntryCount = &entryCount
	}
	if unreadEntryCount != nil {
		unreadCount := int(*unreadEntryCount)
		result.UnreadEntryCount = &unreadCount
	}

	return result
}

func mapToAPIFeedEntry(entry database.Entry, feedEntry database.FeedEntry) api.FeedEntry {
	result := api.FeedEntry{
		Id:          int(entry.ID),
		Link:        entry.Link,
		Title:       entry.Title,
		PublishedAt: entry.PublishedAt.Time,
		HasRead:     feedEntry.HasRead,
		FeedId:      int(feedEntry.FeedID),
	}
	if entry.Author.Valid {
		result.Author = &entry.Author.String
	}
	if entry.Content.Valid {
		result.Content = &entry.Content.String
	}
	if entry.ThumbnailUrl.Valid {
		result.ThumbnailUrl = &entry.ThumbnailUrl.String
	}

	return result
}
