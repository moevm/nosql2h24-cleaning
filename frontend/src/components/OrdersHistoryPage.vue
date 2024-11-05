<template>
  <div class="orders-history">
    <FilterPanel @filter-dates="filterOrders" @reset-filters="resetFilters" />
    <HeaderList title="История заказов" :orders="filteredOrders" />
  </div>
</template>

<script>
import FilterPanel from "../ui/uikit/FilterPanel.vue";
import HeaderList from "../ui/uikit/HeaderList.vue";

export default {
  name: "HistoryOrderPage",
  components: {
    FilterPanel,
    HeaderList
  },
  data() {
    return {
      orders: [
        { id: 1, date: "28.02.2024", time: "10:00", address: "г. Москва, ул. Пушкина, д. 228, кв. 228, п. 7, э. 4", price: 37 },
        { id: 2, date: "01.01.2024", time: "17:00", address: "г. Москва, ул. Пушкина, д. 228, кв. 228, п. 7, э. 4", price: 102 },
        { id: 4, date: "01.01.2024", time: "17:00", address: "г. Москва, ул. Пушкина, д. 228, кв. 228, п. 7, э. 4", price: 102 },
        { id: 5, date: "01.01.2024", time: "17:00", address: "г. Москва, ул. Пушкина, д. 228, кв. 228, п. 7, э. 4", price: 102 },
        { id: 6, date: "01.01.2024", time: "17:00", address: "г. Москва, ул. Пушкина, д. 228, кв. 228, п. 7, э. 4", price: 102 },
        { id: 7, date: "01.01.2024", time: "17:00", address: "г. Москва, ул. Пушкина, д. 228, кв. 228, п. 7, э. 4", price: 102 },
        { id: 8, date: "01.01.2024", time: "17:00", address: "г. Москва, ул. Пушкина, д. 228, кв. 228, п. 7, э. 4", price: 102 },
        { id: 9, date: "01.01.2024", time: "17:00", address: "г. Москва, ул. Пушкина, д. 228, кв. 228, п. 7, э. 4", price: 102 },
        { id: 10, date: "01.01.2024", time: "17:00", address: "г. Москва, ул. Пушкина, д. 228, кв. 228, п. 7, э. 4", price: 102 },
        { id: 11, date: "01.01.2024", time: "17:00", address: "г. Москва, ул. Пушкина, д. 228, кв. 228, п. 7, э. 4", price: 102 },
        { id: 12, date: "15.12.2023", time: "17:00", address: "г. Москва, ул. Пушкина, д. 228, кв. 228, п. 7, э. 4", price: 42 }
      ], // TODO DB request
      filteredOrders: []
    };
  },
  created() {
    this.filteredOrders = this.orders;
  },
  methods: {
    filterOrders(dates) {
      const { startDate, endDate } = dates;
      this.filteredOrders = this.orders.filter(order => {
        const orderDate = new Date(order.date.split('.').reverse().join('-'));
        return (!startDate || new Date(startDate) <= orderDate) &&
               (!endDate || new Date(endDate) >= orderDate);
      });
    },
    resetFilters() {
      this.filteredOrders = this.orders;
    }
  }
};
</script>

<style scoped>
.orders-history {
  padding: 20px;
}
</style>
