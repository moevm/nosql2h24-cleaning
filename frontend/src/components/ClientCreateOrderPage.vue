<template>
  <div>
    <component :is="currentStepComponent" @change-step="changeStep" />
    <ActionButton
      id="next-btn"
      variant="text"
      text="Далее"
      :onClick="changeStep"
    />
  </div>
</template>

<script setup>
import { ref, computed } from 'vue';
import Step1 from './createOrderStep/ContactInfo.vue';
import Step2 from './createOrderStep/AppartmentInfo.vue';
import Step3 from './createOrderStep/ServiceInfo.vue';

const currentStep = ref(1);

const steps = {
  1: Step1,
  2: Step2,
  3: Step3,
};

const currentStepComponent = computed(() => steps[currentStep.value]);

const changeStep = () => {
  if (currentStep >= 3) {
    // create order
    // redirect to order history
  } else {
    currentStep.value += 1;
  }
};
</script>

<style scoped>
.div {
  width: 1000px;
  height: 100%;  
}

#next-btn {
  width: 1000px;
  padding: 12px;
  font-size: 16px;
  background-color: #394cc2;
  color: white;
  border-radius: 8px;
  cursor: pointer;
  border: none;
}

#next-btn:hover {
  background-color: #394cc2;
}
</style>
