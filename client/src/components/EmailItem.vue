<script setup lang="ts">
import type { Email } from '@/services/api'
import { computed, ref } from 'vue'
import MessageIcon from '@/components/icons/MessageIcon.vue'

const maxBodyLength = 180

type ContentProps = {
  isKeyWord: boolean
  text: string
}

const props = defineProps<{
  email: Email
  searchTerm: string
}>()

const isOpen = ref(false)

const toggleOpen = () => {
  isOpen.value = !isOpen.value
}

const completeContent = computed(() => {
  let arr = [{ isKeyWord: false, text: props.email.content }]

  if (!props.searchTerm) {
    return arr
  }

  let j = 0
  while (j < arr.length) {
    let rec = arr[j]
    let record = rec.text.split(props.searchTerm)

    if (record.length > 1) {
      // delete j replace by new
      arr.splice(j, 1)
      let recKeyword = {
        isKeyWord: true,
        text: props.searchTerm,
      }
      for (let k = 0; k < record.length; k++) {
        let r = {
          isKeyWord: false,
          text: record[k],
        }
        if (k == record.length - 1) {
          arr.splice(j + k * 2, 0, r)
        } else {
          arr.splice(j + k * 2, 0, r, recKeyword)
        }
      }
    }

    j = j + record.length
  }

  return arr
})

const displayContent = computed(() => {
  const arr = JSON.parse(
    JSON.stringify(completeContent.value)
  ) as ContentProps[]

  const index = arr.findIndex((item) => item.isKeyWord)

  if (index === -1) {
    arr[0].text = arr[0].text.slice(0, maxBodyLength)
    return arr
  }

  const prev = index === 0 ? null : arr[index - 1]
  const next = index === arr.length - 1 ? null : arr[index + 1]

  if (prev && prev.text.length > maxBodyLength / 2) {
    prev.text = '...' + prev.text.slice(-maxBodyLength / 2)
  }

  if (next && next.text.length > maxBodyLength / 2) {
    next.text = next.text.slice(0, maxBodyLength / 2) + '...'
  }

  return [prev, arr[index], next].filter(Boolean) as ContentProps[]
})

const date = computed(() => {
  const date = new Date(props.email.date)
  return date.toLocaleDateString()
})

const fromUser = computed(() => props.email.from.split('@')[0])
</script>

<template>
  <li class="item">
    <i><MessageIcon /></i>
    <div class="item-container">
      <div :class="{ details: true, isOpen }" @click="toggleOpen">
        <h3>{{ email.subject }}</h3>
        <p>
          <span v-for="item in displayContent" :key="item.text" v-bind="item">
            <span v-if="item.isKeyWord" class="highlight">{{ item.text }}</span>
            <span v-else>{{ item.text }}</span>
          </span>
        </p>

        <div class="contact">
          <span>👤 {{ fromUser }}</span>
          <span>📅 {{ date }}</span>
        </div>
      </div>
      <Transition name="fade" :duration="300">
        <div class="json" v-if="isOpen">
          <span v-for="item in completeContent" :key="item.text" v-bind="item">
            <span v-if="item.isKeyWord" class="highlight">{{ item.text }}</span>
            <span v-else>{{ item.text }}</span>
          </span>
        </div>
      </Transition>
    </div>
  </li>
</template>

<style scoped>
.item {
  margin-top: 2rem;
  display: flex;
}

.item-container {
  display: flex;
  flex-direction: column;
}

.details {
  flex: 1;
  margin-left: 1rem;
}

.isOpen {
  background-color: var(--color-background-soft);
}

.contact {
  gap: 1.5rem;
  display: flex;
  justify-content: left;
  margin-top: 0.3rem;
}

.json {
  padding: 1rem;
  transition: all 2s ease;
}

.highlight {
  color: hsla(160, 100%, 37%, 1);
  background-color: hsla(160, 100%, 37%, 0.2);
}

.fade-enter-active,
.fade-leave-active {
  transition: all 0.5s ease-out;
  overflow: hidden;
}

.fade-enter-from,
.fade-leave-to {
  transform: translateY(-20px);
  opacity: 0;
}

i {
  display: flex;
  place-items: center;
  place-content: center;
  width: 32px;
  height: 32px;

  color: var(--color-text);
}

h3 {
  font-size: 1.2rem;
  font-weight: 500;
  margin-bottom: 0.4rem;
  color: var(--color-heading);
}

@media (min-width: 1024px) {
  .item {
    margin-top: 0;
    padding: 0.4rem 0 0.4rem calc(var(--section-gap) / 3);
  }

  .details {
    margin: 0;
    padding: 1rem;
    border-radius: 1rem;
  }

  .details:hover {
    cursor: pointer;
    background-color: var(--color-background-soft);
  }

  i {
    top: calc(50% - 25px);
    left: -26px;
    position: absolute;
    border: 1px solid var(--color-border);
    background: var(--color-background);
    border-radius: 8px;
    width: 50px;
    height: 50px;
  }

  .item:before {
    content: ' ';
    border-left: 1px solid var(--color-border);
    position: absolute;
    left: 0;
    bottom: calc(50% + 25px);
    height: calc(50% - 25px);
  }

  .item:after {
    content: ' ';
    border-left: 1px solid var(--color-border);
    position: absolute;
    left: 0;
    top: calc(50% + 25px);
    height: calc(50% - 25px);
  }

  .item:first-of-type:before {
    display: none;
  }

  .item:last-of-type:after {
    display: none;
  }
}
</style>
