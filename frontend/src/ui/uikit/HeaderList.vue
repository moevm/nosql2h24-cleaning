<template>
  <div class="header-list">
    <h1>{{ title }}</h1>
    <div class="orders-list">
      <OrderItem v-for="order in orders" :key="order.id" :order="order" />
    </div>
  </div>
</template>

<script>
import OrderItem from "./OrderItem.vue";

export default {
  name: "HeaderList",
  components: {
    OrderItem
  },
  props: {
    title: {
      type: String,
      required: true
    },
    orders: {
      type: Array,
      required: true
    }
  },
};
</script>

<style scoped>
.header-list {
  display: flex;
  flex-direction: column;
  height: 90vh;
  box-sizing: border-box;
  padding-left: 10px;
  padding-right: 10px;
  border: 2px solid grey;
  border-radius: 25px;
  text-align: center;
}

.title {
  flex-shrink: 0;
  font-size: 24px;
  margin-bottom: 10px;
}

.orders-list {
  display: grid;
  gap: 10px;
  flex-grow: 1;
  overflow-y: scroll;
  scrollbar-width: none;
}

.orders-list::-webkit-scrollbar {
  width: 0px;
  background: transparent;
}
</style>