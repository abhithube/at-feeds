/* prettier-ignore-start */

/* eslint-disable */

// @ts-nocheck

// noinspection JSUnusedGlobalSymbols

// This file is auto-generated by TanStack Router

// Import Routes

import { Route as rootRoute } from './routes/__root'
import { Route as SavedImport } from './routes/saved'
import { Route as IndexImport } from './routes/index'
import { Route as FeedsFeedIdImport } from './routes/feeds.$feedId'

// Create/Update Routes

const SavedRoute = SavedImport.update({
  path: '/saved',
  getParentRoute: () => rootRoute,
} as any)

const IndexRoute = IndexImport.update({
  path: '/',
  getParentRoute: () => rootRoute,
} as any)

const FeedsFeedIdRoute = FeedsFeedIdImport.update({
  path: '/feeds/$feedId',
  getParentRoute: () => rootRoute,
} as any)

// Populate the FileRoutesByPath interface

declare module '@tanstack/react-router' {
  interface FileRoutesByPath {
    '/': {
      preLoaderRoute: typeof IndexImport
      parentRoute: typeof rootRoute
    }
    '/saved': {
      preLoaderRoute: typeof SavedImport
      parentRoute: typeof rootRoute
    }
    '/feeds/$feedId': {
      preLoaderRoute: typeof FeedsFeedIdImport
      parentRoute: typeof rootRoute
    }
  }
}

// Create and export the route tree

export const routeTree = rootRoute.addChildren([
  IndexRoute,
  SavedRoute,
  FeedsFeedIdRoute,
])

/* prettier-ignore-end */
