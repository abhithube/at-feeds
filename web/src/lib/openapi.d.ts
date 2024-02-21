/**
 * This file was auto-generated by openapi-typescript.
 * Do not make direct changes to the file.
 */


export interface paths {
  "/feeds": {
    get: operations["listFeeds"];
    post: operations["createFeed"];
  };
  "/feeds/{id}": {
    get: operations["getFeed"];
    delete: operations["deleteFeed"];
  };
  "/feeds/import": {
    post: operations["importFeeds"];
  };
  "/feeds/export": {
    post: operations["exportFeeds"];
  };
  "/entries": {
    get: operations["listEntries"];
  };
  "/entries/{id}": {
    patch: operations["updateEntry"];
  };
}

export type webhooks = Record<string, never>;

export interface components {
  schemas: {
    Feed: {
      id: number;
      /** Format: uri */
      url?: string | null;
      /** Format: uri */
      link: string;
      title: string;
      entryCount?: number;
      unreadCount: number;
    };
    Entry: {
      id: number;
      /** Format: uri */
      link: string;
      title: string;
      /** Format: date-time */
      publishedAt: string;
      author: string | null;
      content: string | null;
      /** Format: uri */
      thumbnailUrl: string | null;
      hasRead: boolean;
    };
    CreateFeed: {
      /** Format: uri */
      url: string;
    };
    UpdateEntry: {
      hasRead: boolean;
    };
    /** Format: binary */
    File: string;
    Error: {
      message: string;
    };
  };
  responses: never;
  parameters: never;
  requestBodies: never;
  headers: never;
  pathItems: never;
}

export type $defs = Record<string, never>;

export type external = Record<string, never>;

export interface operations {

  listFeeds: {
    parameters: {
      query?: {
        limit?: number;
        page?: number;
      };
    };
    responses: {
      /** @description Paginated feed list response */
      200: {
        content: {
          "application/json": {
            hasMore: boolean;
            data: components["schemas"]["Feed"][];
          };
        };
      };
    };
  };
  createFeed: {
    requestBody: {
      content: {
        "application/json": components["schemas"]["CreateFeed"];
      };
    };
    responses: {
      /** @description Feed response */
      201: {
        content: {
          "application/json": components["schemas"]["Feed"];
        };
      };
      /** @description Invalid feed response */
      400: {
        content: {
          "application/json": components["schemas"]["Error"];
        };
      };
    };
  };
  getFeed: {
    parameters: {
      path: {
        id: number;
      };
    };
    responses: {
      /** @description Feed response */
      200: {
        content: {
          "application/json": components["schemas"]["Feed"];
        };
      };
      /** @description Feed not found */
      404: {
        content: {
          "application/json": components["schemas"]["Error"];
        };
      };
    };
  };
  deleteFeed: {
    parameters: {
      path: {
        id: number;
      };
    };
    responses: {
      /** @description Successful response */
      204: {
        content: never;
      };
      /** @description Feed not found */
      404: {
        content: {
          "application/json": components["schemas"]["Error"];
        };
      };
    };
  };
  importFeeds: {
    requestBody?: {
      content: {
        "application/octet-stream": components["schemas"]["File"];
      };
    };
    responses: {
      /** @description Empty response */
      200: {
        content: never;
      };
      /** @description Internal error response */
      500: {
        content: {
          "application/json": components["schemas"]["Error"];
        };
      };
    };
  };
  exportFeeds: {
    responses: {
      /** @description OPML file download */
      200: {
        content: {
          "application/octet-stream": components["schemas"]["File"];
        };
      };
      /** @description Internal error response */
      500: {
        content: {
          "application/json": components["schemas"]["Error"];
        };
      };
    };
  };
  listEntries: {
    parameters: {
      query?: {
        feedId?: number;
        hasRead?: boolean;
        limit?: number;
        page?: number;
      };
    };
    responses: {
      /** @description Paginated entry list response */
      200: {
        content: {
          "application/json": {
            hasMore: boolean;
            data: components["schemas"]["Entry"][];
          };
        };
      };
    };
  };
  updateEntry: {
    parameters: {
      path: {
        id: number;
      };
    };
    requestBody: {
      content: {
        "application/json": components["schemas"]["UpdateEntry"];
      };
    };
    responses: {
      /** @description Updated entry */
      200: {
        content: {
          "application/json": components["schemas"]["Entry"];
        };
      };
      /** @description Entry not found */
      404: {
        content: {
          "application/json": components["schemas"]["Error"];
        };
      };
    };
  };
}
