import { Button } from '@/components/ui/button'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog'
import {
  Form,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
} from '@/components/ui/form'
import { Input } from '@/components/ui/input'
import { client } from '@/lib/client'
import { zodResolver } from '@hookform/resolvers/zod'
import { useMutation, useQueryClient } from '@tanstack/react-query'
import { useNavigate } from '@tanstack/react-router'
import { Loader2 } from 'lucide-react'
import { useState } from 'react'
import { useForm } from 'react-hook-form'
import { z } from 'zod'

const formSchema = z.object({
  url: z.string().url(),
})

export function AddFeedModal() {
  const [open, setOpen] = useState(true)
  const [loading, setLoading] = useState(false)

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      url: '',
    },
  })

  const navigate = useNavigate()
  const queryClient = useQueryClient()

  const { mutateAsync: addFeed } = useMutation({
    mutationFn: async (values: z.infer<typeof formSchema>) => {
      const res = await client.POST('/feeds', {
        body: {
          url: values.url,
        },
      })
      if (res.error) {
        throw new Error(res.error.message)
      }

      return res.data
    },
    onMutate: () => {
      setLoading(true)
    },
    onSuccess: (data) => {
      setLoading(false)
      setOpen(false)

      queryClient.invalidateQueries({
        predicate: (query) => query.queryKey[0] === '/feeds',
      })

      navigate({
        to: '/feeds/$feedId',
        params: {
          feedId: data.id.toString(),
        },
      })
    },
  })

  return (
    <Dialog
      open={open}
      onOpenChange={(open) => {
        if (!open) {
          navigate({
            params: {},
            search: {},
          })
        }
      }}
    >
      <DialogContent className="max-w-[425px]">
        <Form {...form}>
          <form onSubmit={form.handleSubmit((data) => addFeed(data))}>
            <DialogHeader>
              <DialogTitle>Add Feed</DialogTitle>
              <DialogDescription>
                Subscribe to a RSS or Atom feed and receive the latest updates.
              </DialogDescription>
            </DialogHeader>
            <div className="flex flex-col py-4">
              <FormField
                control={form.control}
                name="url"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>URL</FormLabel>
                    <Input {...field} />
                    <FormDescription>
                      URL of the RSS or Atom Feed
                    </FormDescription>
                  </FormItem>
                )}
              ></FormField>
            </div>
            <DialogFooter>
              <Button disabled={loading}>
                {loading && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
                Submit
              </Button>
            </DialogFooter>
          </form>
        </Form>
      </DialogContent>
    </Dialog>
  )
}
