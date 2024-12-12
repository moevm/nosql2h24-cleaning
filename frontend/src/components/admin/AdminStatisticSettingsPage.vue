<script setup lang="ts">
import { ref, Ref } from 'vue'
import MainContainer from '../../ui/uikit/containers/MainContainer.vue'
import Dialog from '../../ui/uikit/Dialog.vue'
import PanelContainer from '../../ui/uikit/containers/PanelContainer.vue'
const isDialogImportVisible: Ref<boolean> = ref(false)
const isDialogExportVisible: Ref<boolean> = ref(false)
const isDialogFiltersVisible: Ref<boolean> = ref(false)

const isLoading = ref(false)
const loadingProgress = ref(0)
let loadingInterval: number

function openDialog(type: string): void {
	switch (type) {
		case 'IMPORT': 
			isDialogImportVisible.value = true
			break
		case 'EXPORT':
			isDialogExportVisible.value = true
			break;
		case 'FILTERS':
			isDialogFiltersVisible.value = true
			break;
	}
}

function closeDialog(): void {
  isDialogImportVisible.value = false
	isDialogExportVisible.value = false
	isDialogFiltersVisible.value = false
}

function handleFileDrop(event: DragEvent) {
  const files = event.dataTransfer?.files
  if (files?.length) {
    console.log('Файл добавлен для загрузки:', files[0])
  }
}

function startImport() {
  isLoading.value = true
  loadingProgress.value = 0
  loadingInterval = setInterval(() => {
    if (loadingProgress.value < 100) {
      loadingProgress.value += 1
    } else {
      completeImport()
    }
  }, 100)
}

function completeImport() {
  clearInterval(loadingInterval)
  isLoading.value = false
  closeDialog()
  console.log('Импорт данных закончен')
}

function cancelImport() {
  clearInterval(loadingInterval)
  resetLoading()
  console.log('Импорт данных прерван')
}

function resetLoading() {
  isLoading.value = false
  loadingProgress.value = 0
}

function fetchExportData() {
	console.log('Экспорт данных');
}

</script>

<template>
  <MainContainer>
    <PanelContainer
      height="10%"
      width="95%"
    >
      <ActionButton
        text="Фильтры"
        type="add"
        color="#394cc2"
        @click="openDialog('FILTERS')"
      ></ActionButton>
      <ActionButton
        text="Импорт"
        type="add"
        color="#394cc2"
        @click="openDialog('IMPORT')"
      ></ActionButton>
      <ActionButton
        text="Экспорт"
        type="add"
        color="#394cc2"
        @click="openDialog('EXPORT')"
      ></ActionButton>
    </PanelContainer>

    <Dialog
      title="Экспорт"
      :visible="isDialogExportVisible"
      dialogMaxWidth="30%"
      @update:visible="closeDialog"
			>
			<div>
        <p>Вы уверены, что хотите экспортировать все данные сервиса?</p>
        <div class="button-container">
          <button class="no-button" @click="closeDialog">Нет</button>
          <button class="yes-button" @click="fetchExportData">Да</button>
        </div>
      </div>
    </Dialog>
		<Dialog
      title="Импорт"
      :visible="isDialogImportVisible"
      dialogMaxWidth="30%"
      @update:visible="closeDialog"
			>
			<div>
        <p>Вы уверены, что хотите загрузить новые данные? Это удалит все текущие данные.</p>

        <div
          class="drag-drop-area"
          @dragover.prevent
          @drop.prevent
          @drop="handleFileDrop"
        >
          Перетащите файл для загрузки
        </div>

        <div v-if="isLoading">
          <p>Загрузка... Подождите или нажмите 'Стоп', чтобы отменить.</p>
          <progress :value="loadingProgress" max="100"></progress>
          <button class="stop-button" @click="cancelImport">Стоп</button>
        </div>

        <div class="button-container" v-else>
          <button class="no-button" @click="closeDialog">Нет</button>
          <button class="yes-button" @click="startImport()">Да</button>
        </div>
      </div>
    </Dialog>
		<Dialog
      title="Фильтры"
      :visible="isDialogFiltersVisible"
      dialogMaxWidth="30%"
      @update:visible="closeDialog"
			>
			
    </Dialog>
  </MainContainer>
</template>

<style>
.drag-drop-area {
  border: 2px dashed #ccc;
  padding: 20px;
  text-align: center;
  margin: 20px 0;
  color: #666;
}

.no-button {
  background-color: red;
  color: white;
  padding: 10px 20px;
  border: none;
  cursor: pointer;
}

.yes-button {
  background-color: green;
  color: white;
  padding: 10px 20px;
  border: none;
  cursor: pointer;
}

.stop-button {
  background-color: orange;
  color: white;
  padding: 10px 20px;
  border: none;
  cursor: pointer;
}

progress {
  width: 100%;
  margin-top: 10px;
}
</style>