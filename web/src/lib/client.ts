import createClient from 'openapi-fetch'
import { paths } from './openapi'

export const client = createClient<paths>({
  baseUrl: import.meta.env.VITE_BACKEND_URL,
})

client.use({
  onResponse: async (res) => {
    if (res.status >= 400) {
      const body = await res.clone().json()

      throw new Error(body.message)
    }

    return undefined
  },
})
