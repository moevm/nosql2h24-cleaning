<template>
  <div class="filter-panel">
    <ActionButton
      id="list-filters-btn"
      :text="'Все'"
    ></ActionButton>
    <DateFilter @filter-dates="updateDateFilter" />
    <ActionButton
      @click="updateDateFilter"
      id="search-btn"
      :text="'Поиск'"
    ></ActionButton>
  </div>
</template>

<script>
import ActionButton from "./ActionButton.vue";
import DateFilter from "./DateFilter.vue";

export default {
  name: "FilterPanel",
  components: { ActionButton, DateFilter },
  methods: {
    updateDateFilter(dates) {
      this.$emit("filter-dates", dates);
    }
  }
};
</script>

<style scoped>
.filter-panel {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 10px;
  margin-bottom: 10px;
}
.btn {
  font-size: 16px;
  background-color: #3846c0;
  color: black;
  padding-left: 20px;
  padding-right: 20px;
  height: 40px;
}
</style>
