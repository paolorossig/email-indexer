<script setup lang="ts">
import type { Email } from '@/services/api'
import EmailItem from '@/components/EmailItem.vue'

defineProps<{
  emails: Email[] | null
  searchTerm: string
  loading: boolean
}>()
</script>

<template>
  <h1 v-if="!emails?.length && loading" class="center">Loading Emails..</h1>
  <h1 v-else-if="!emails?.length" class="center">No Emails found!</h1>
  <ul v-else>
    <EmailItem
      v-for="email in emails"
      :key="email.message_id"
      :email="email"
      :searchTerm="searchTerm"
    />
  </ul>
</template>

<style scoped>
.center {
  display: grid;
  place-items: center;
  padding: 4rem 0;
}
</style>
