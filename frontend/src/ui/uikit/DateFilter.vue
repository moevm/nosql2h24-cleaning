<template>
  <div class="date-filter">
    <input type="date" v-model="startDate" placeholder="C" @change="updateDateFilter" />
    <input type="date" v-model="endDate" placeholder="По" @change="updateDateFilter" />
  </div>
</template>

<script>
import ActionButton from "./ActionButton.vue";

export default {
  name: "DateFilter",
  components: { ActionButton },
  data() {
    return {
      startDate: null,
      endDate: null
    };
  },
  methods: {
    updateDateFilter() {
      this.$emit("filter-dates", { startDate: this.startDate, endDate: this.endDate });
    }
  }
};
</script>

<style scoped>
.date-filter {
  border: 2px solid grey;
  border-radius: 25px;
  padding: 5px;
  padding-left: 10px;
}
</style>
