import { Checkbox } from '@/components/ui/checkbox'
import { client } from '@/lib/client'
import { Entry, UpdateEntryBody } from '@/lib/types'
import { useMutation } from '@tanstack/react-query'
import { useState } from 'react'

type EntryCardActionsProps = {
  entry: Entry
}

export function EntryCardActions({ entry }: EntryCardActionsProps) {
  const [hasRead, setRead] = useState(entry.hasRead)

  const { mutateAsync: updateEntry } = useMutation({
    mutationFn: async (variables: UpdateEntryBody & { id: number }) => {
      const res = await client.PATCH('/entries/{id}', {
        params: {
          path: {
            id: +variables.id,
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
            id: +entry.id,
            hasRead: !checked,
          })
        }
      />
    </div>
  )
}
