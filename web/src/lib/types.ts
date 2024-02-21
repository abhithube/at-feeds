import { components, operations } from './openapi'

type schema = components['schemas']

export type Feed = schema['Feed']
export type Entry = schema['Entry']

export type ListFeedsQuery = Omit<
  NonNullable<operations['listFeeds']['parameters']['query']>,
  'page'
>

export type ListEntriesQuery = Omit<
  NonNullable<operations['listEntries']['parameters']['query']>,
  'page'
>

export type UpdateEntryBody = NonNullable<
  operations['updateEntry']['requestBody']['content']['application/json']
>
