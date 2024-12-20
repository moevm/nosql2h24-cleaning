<script setup lang="ts">
import { Ref, ref, computed } from 'vue'
import MainContainer from '../../ui/uikit/containers/MainContainer.vue'
import ActionButton from '../../ui/uikit/ActionButton.vue'
import Step1 from '../createOrderStep/ContactInfo.vue'
import Step2 from '../createOrderStep/AppartmentInfo.vue'
import Step3 from '../createOrderStep/ServiceInfo.vue'

const currentStep: Ref<number> = ref(1)
const steps: any = {
  1: Step1,
  2: Step2,
  3: Step3,
}
const currentStepComponent = computed(() => steps[currentStep.value])
const currentPrice = ref<number>(0)
const isFormValid = ref<boolean>(false)

function changeStep(): void {
  if (currentStep.value >= 3) {
    // create order
    
    // redirect to order history
  } else {
    currentStep.value += 1;
    isFormValid.value = false
  }
}

function updateFormValidity(isValid: boolean): void {
  isFormValid.value = isValid
}
function updatePrice(num: number): void {
  currentPrice.value = num
}

</script>

<template>
  <MainContainer
    justifyContent="flex-start"
  >
    <component
      :is="currentStepComponent"
      @update-form-validity="updateFormValidity"
      @update-price="updatePrice"
      :currentPrice="currentPrice"
    ></component>
    <ActionButton
      id="next-btn"
      text="Далее"
      type="next"
      variant="flat"
      color="#394cc2"
      :onClick="changeStep"
      :disabled="!isFormValid"
    ></ActionButton>
  </MainContainer>
</template>

<style scoped>
#next-btn {
  width: 95%;
}
</style>
