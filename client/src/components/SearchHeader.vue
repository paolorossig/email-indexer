<script setup lang="ts">
import SearchIcon from '@/components/icons/SearchIcon.vue'
</script>

<template>
  <header>
    <img
      alt="Vue logo"
      class="logo"
      src="@/assets/email.svg"
      width="125"
      height="125"
    />
    <div class="wrapper">
      <div class="greetings">
        <h1 class="green">Deep search</h1>
        <h3>
          Blazingly fast email searcher powered with
          <a href="https://zincsearch.com/" target="_blank" rel="noopener"
            >Zinc</a
          >
          + <a href="https://go.dev/" target="_blank" rel="noopener">Go</a>.
          Start searching!
        </h3>
      </div>
      <form class="search">
        <input type="text" name="term" placeholder="Search a word..." />
        <button type="submit">
          <i>
            <SearchIcon />
          </i>
          <!-- <img src="@/assets/search-mini.svg" width="25" height="25" /> -->
        </button>
      </form>
    </div>
  </header>
</template>

<style scoped>
i {
  display: flex;
  place-items: center;
  place-content: center;
  width: 32px;
  height: 32px;

  color: var(--color-text);
}

header {
  line-height: 1.5;
  max-height: 100vh;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
  background-color: #00bd7e;
  border-radius: 1rem;
}

@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }
}

h1 {
  font-weight: 500;
  font-size: 2.6rem;
  top: -10px;
}

h3 {
  font-size: 1.2rem;
}

input {
  width: 100%;
  height: 40px;
  border: 1px solid var(--color-border);
  border-radius: 8px;
  padding: 0 10px;
  font-size: 1rem;
}

input:focus {
  outline: none;
  border-color: #00bd7e;
}

.search {
  width: 100%;
  margin-top: 20px;
  position: relative;
}

.search button {
  position: absolute;
  right: 0;
  top: 0;
  height: 40px;
  width: 40px;
  border: 0;
  background-color: #00bd7e;
  border-radius: 8px;
  cursor: pointer;
  transition: background-color 0.3s ease-in-out;
  display: flex;
  justify-content: center;
  align-items: center;
}

.search button:hover {
  background-color: hsl(160, 100%, 30%);
}

.greetings h1,
.greetings h3 {
  text-align: center;
}

@media (min-width: 1024px) {
  .greetings h1,
  .greetings h3 {
    text-align: left;
  }
}
</style>
