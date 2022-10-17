import { defineComponent, ref } from 'vue'
import SearchHeader from '@/components/SearchHeader.vue'
import SearchResults from '@/components/SearchResults.vue'

import { searchInEmails, type ApiError, type Email } from '../services/api'
import axios, { AxiosError } from 'axios'

interface InputEvent extends Event {
  target: HTMLInputElement & { term: { value: string } }
}

export default defineComponent({
  name: 'HomePage',
  components: {
    SearchHeader,
    SearchResults,
  },
  setup() {
    const emails = ref<Email[]>([])
    const loading = ref(false)
    const searchTerm = ref('')

    const getEmails = async () => {
      loading.value = true

      try {
        const response = await searchInEmails(searchTerm.value)
        emails.value = response.data.emails
      } catch (error) {
        if (axios.isAxiosError(error)) {
          const serverError = error as AxiosError<ApiError>
          console.error('API Error -', serverError.response?.data.error)
        } else {
          console.error('Unknown Error -', error)
        }
      } finally {
        loading.value = false
      }
    }

    const searchNewTerm = (e: InputEvent) => {
      searchTerm.value = e.target.term.value
      getEmails()
    }

    getEmails()

    return {
      emails,
      loading,
      searchTerm,
      getEmails,
      searchNewTerm,
    }
  },
})
