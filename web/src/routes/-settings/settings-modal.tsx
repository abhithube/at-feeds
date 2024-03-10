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
import { zfd } from 'zod-form-data'

const formSchema = z.object({
  file: zfd.file(),
})

export function SettingsModal() {
  const [open, setOpen] = useState(true)
  const [loading, setLoading] = useState(false)

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
  })

  const navigate = useNavigate()
  const queryClient = useQueryClient()

  const { mutateAsync: importBackup } = useMutation({
    mutationFn: async (values: z.infer<typeof formSchema>) => {
      const res = await client.POST('/feeds/import', {
        body: values.file as any,
        bodySerializer: (body) => {
          const fd = new FormData()
          fd.append('file', body!)

          return fd
        },
      })
      if (res.error) {
        throw new Error(res.error.message)
      }
    },
    onMutate: () => {
      setLoading(true)
    },
    onSuccess: () => {
      setLoading(false)
      setOpen(false)

      queryClient.invalidateQueries({
        predicate: (query) =>
          ['/feeds', '/entries'].includes(query.queryKey[0] as any),
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
          <form onSubmit={form.handleSubmit((data) => importBackup(data))}>
            <DialogHeader>
              <DialogTitle>Import Feeds</DialogTitle>
              <DialogDescription>
                Upload an OPML file to import feeds.
              </DialogDescription>
            </DialogHeader>
            <div className="flex flex-col py-4">
              <FormField
                control={form.control}
                name="file"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>File</FormLabel>
                    <Input
                      type="file"
                      onChange={(ev) =>
                        field.onChange(
                          ev.target.files ? ev.target.files[0] : null,
                        )
                      }
                    />
                    <FormDescription>OPML file to upload</FormDescription>
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
