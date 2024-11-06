<script setup lang="ts">

</script>

<template>
  <div class="container">
    <div class="my-addresses-container">
      <h1>
        Мои адреса
      </h1>
      <div class="addresses-container">
        <div class="address-card">
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.container {
  max-width: 100%;
  max-height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 100%;
  width: 80%;
  background-color: white;
  gap: 20px;
  padding-top: 20px;
  padding-bottom: 20px;
}
.my-addresses-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  height: 100%;
  width: 95%;
  background-color: white;
  border: 2px solid gray;
  border-radius: 20px;
}
.addresses-container {
  display: flex;
  flex-direction: row;
  justify-content: space-evenly;
  align-content: flex-start;
  flex-wrap: wrap;
  height: 100%;
  width: 100%;
  gap: 10px;
  padding-left: 10px;
  padding-right: 10px;
}
.address-card {
  display: flex;
  flex-direction: row;
  justify-content: top;
  align-items: center;
  height: 20%;
  width: 49%;
  border: 3px solid #394cc2;
  border-radius: 20px;
}
h1 {
  font-size: 42px;
  font-weight: bold;
}
</style>
