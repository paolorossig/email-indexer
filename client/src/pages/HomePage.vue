<script src="./HomePage.ts" />

<template>
  <SearchHeader @submit.prevent="searchNewTerm" />
  <SearchResults :emails="emails" :searchTerm="searchTerm" :loading="loading" />
</template>
