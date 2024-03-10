import { components, operations } from './openapi'

type schema = components['schemas']

export type Collection = schema['Collection']
export type Feed = schema['Feed']
export type Entry = schema['FeedEntry']

export type ListCollectionsQuery = Omit<
  NonNullable<operations['listCollections']['parameters']['query']>,
  'page'
>

export type ListFeedsQuery = Omit<
  NonNullable<operations['listFeeds']['parameters']['query']>,
  'page'
>

export type ListFeedEntriesQuery = Omit<
  NonNullable<operations['listFeedEntries']['parameters']['query']>,
  'page'
>

export type UpdateFeedEntryBody = NonNullable<
  operations['updateFeedEntry']['requestBody']['content']['application/json']
>
