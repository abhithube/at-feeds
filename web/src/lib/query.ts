import {
  QueryClient,
  UseInfiniteQueryOptions,
  infiniteQueryOptions,
  queryOptions,
} from '@tanstack/react-query'
import { client } from './client'
import {
  ListCollectionsQuery,
  ListFeedEntriesQuery,
  ListFeedsQuery,
} from './types'

export async function ensureInfiniteQueryData(
  queryClient: QueryClient,
  options: UseInfiniteQueryOptions,
) {
  const data = queryClient.getQueryData(options.queryKey)
  if (!data) {
    await queryClient.fetchInfiniteQuery(options)
  }
}

export const collectionsQueryOptions = (query: ListCollectionsQuery) =>
  queryOptions({
    queryKey: ['/collections', query],
    queryFn: async () => {
      const res = await client.GET('/collections', {
        params: {
          query,
        },
      })

      return res.data!
    },
  })

export const feedsQueryOptions = (query: ListFeedsQuery) =>
  queryOptions({
    queryKey: ['/feeds', query],
    queryFn: async () => {
      const res = await client.GET('/feeds', {
        params: {
          query,
        },
      })

      return res.data!
    },
  })

export const feedQueryOptions = (id: number) =>
  queryOptions({
    queryKey: ['/feeds', id],
    queryFn: async () => {
      const res = await client.GET('/feeds/{id}', {
        params: {
          path: {
            id,
          },
        },
      })

      return res.data!
    },
  })

export const feedEntriesQueryOptions = (query: ListFeedEntriesQuery) =>
  infiniteQueryOptions({
    queryKey: ['/entries', query],
    queryFn: async ({ pageParam }) => {
      const res = await client.GET('/entries', {
        params: {
          query: {
            ...query,
            page: pageParam,
          },
        },
      })

      return res.data!
    },
    initialPageParam: 1,
    getNextPageParam: (lastPage, __, lastPageParam) => {
      return lastPage.hasMore ? lastPageParam + 1 : undefined
    },
  })
