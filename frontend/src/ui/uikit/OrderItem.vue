<template>
  <div class="order-item">
    <p>Заказ {{ order.date }} в {{ order.time }}</p>
    <p>{{ order.address }}</p>
    <p>{{ order.price }}$</p>
  </div>
</template>

<script>
export default {
  name: "OrderItem",
  props: {
    order: {
      type: Object,
      required: true
    }
  }
};
</script>

<style scoped>
.order-item {
  border: 3px solid #3846c0;
  border-radius: 15px;
  padding: 10px;
  text-align: left;
}
</style>
