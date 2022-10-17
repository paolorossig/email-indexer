import axios from 'axios'

export type Email = {
  message_id: string
  subject: string
  date: string
  from: string
  to: string
  content: string
  filepath: string
}

export type ApiError = { status: number; error: string }

export const api = axios.create({
  baseURL: 'http://localhost:8080/api',
})

type EmailsSearchResult = { emails: Email[] }

export const searchInEmails = (query: string) => {
  let term = query
  if (!query) term = 'email'

  return api.get<EmailsSearchResult>(`/emails/search?q=${term}`)
}
