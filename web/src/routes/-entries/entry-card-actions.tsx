import { Checkbox } from '@/components/ui/checkbox'
import { client } from '@/lib/client'
import { Entry, UpdateFeedEntryBody } from '@/lib/types'
import { useMutation } from '@tanstack/react-query'
import { useState } from 'react'

type EntryCardActionsProps = {
  entry: Entry
}

export function EntryCardActions({ entry }: EntryCardActionsProps) {
  const [hasRead, setRead] = useState(entry.hasRead)

  const { mutateAsync: updateEntry } = useMutation({
    mutationFn: async (variables: UpdateFeedEntryBody) => {
      const res = await client.PATCH('/feeds/{feedId}/entries/{entryId}', {
        params: {
          path: {
            feedId: entry.feedId,
            entryId: entry.id,
          },
        },
        body: {
          hasRead: variables.hasRead,
        },
      })
      if (res.error) {
        throw new Error(res.error.message)
      }

      return res.data
    },
    onSuccess: (data) => {
      setRead(data.hasRead)
    },
  })

  return (
    <div className="flex items-end">
      <Checkbox
        checked={!hasRead}
        onCheckedChange={(checked) =>
          updateEntry({
            hasRead: !checked,
          })
        }
      />
    </div>
  )
}
