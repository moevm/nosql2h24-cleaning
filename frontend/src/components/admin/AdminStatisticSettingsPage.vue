<script setup lang="ts">
import { ref, Ref } from 'vue'
import { useRouter } from 'vue-router'
import MainContainer from '../../ui/uikit/containers/MainContainer.vue'
import Dialog from '../../ui/uikit/Dialog.vue'
import PanelContainer from '../../ui/uikit/containers/PanelContainer.vue'
import ActionButton from '../../ui/uikit/ActionButton.vue'
import { exportDumps, uploadDumps } from '../../api/request'

const router = useRouter()

const isDialogImportVisible: Ref<boolean> = ref(false)
const isDialogExportVisible: Ref<boolean> = ref(false)
const isDialogFiltersVisible: Ref<boolean> = ref(false)

const isLoading = ref(false)
const loadingProgress = ref(0)
const remainingSeconds = ref(10)
let loadingInterval: number
let selectedFile: File | null = null
const fileName = ref<string | null>(null)

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

function handleCancelAndReset() {
  cancelImport()
  resetLoading()
  resetFile()
  closeDialog()
}

function handleFileDrop(event: DragEvent) {
  const files = event.dataTransfer?.files
  if (files?.length) {
    const file = files[0]
    if (file.type === 'application/json') {
      selectedFile = file
      fileName.value = file.name
    } else {
      console.error('Файл должен быть в формате JSON')
      alert('Пожалуйста, выберите файл формата JSON')
    }
  }
}

function confirmImportFile() {
  if (selectedFile) {
    startImport()
  } else {
    console.error('Нет выбранного файла')
    alert('Пожалуйста, сначала выберите файл JSON')
  }
}

function startImport() {
  isLoading.value = true;
  loadingProgress.value = 0;
  remainingSeconds.value = 10;

  let totalDuration = 10000;
  let intervalDuration = 100;
  let progressIncrement = (100 / (totalDuration / intervalDuration));

  loadingInterval = setInterval(() => {
    if (loadingProgress.value < 100) {
      loadingProgress.value += progressIncrement;

      if (remainingSeconds.value > 0 && loadingProgress.value % 10 === 0) {
        remainingSeconds.value -= 1;
      }
    } else {
      completeImport();
    }
  }, intervalDuration);
}

function completeImport() {
  clearInterval(loadingInterval)
  isLoading.value = false
  loadingProgress.value = 0
  remainingSeconds.value = 10
  closeDialog()
  if (selectedFile) {
    uploadFile(selectedFile)
    resetFile()
  }
  console.log('Загрузка завершена')
}

function cancelImport() {
  clearInterval(loadingInterval)
  resetLoading()
  console.log('Импорт данных прерван')
}

function resetLoading() {
  clearInterval(loadingInterval)
  isLoading.value = false
  loadingProgress.value = 0
  remainingSeconds.value = 10
}

function resetFile() {
  selectedFile = null
  fileName.value = null
}

function uploadFile(file: File) {
  uploadDumps(file)
    .then(() => {
      closeDialog();
      router.push('/login')
    })
    .catch((error) => {
      console.error('Ошибка при загрузке файла:', error);
      alert('Произошла ошибка при загрузке файла. Возможно он не соответсвует формату json');
      resetLoading()
      resetFile()
    });
}

function fetchExportData() {
  exportDumps()
    .then(() => {
      console.log('Данные успешно экспортированы.');
    })
    .catch((error) => {
      console.error('Ошибка экспорта данных:', error);
      alert('Произошла ошибка при экспорте данных. Пожалуйста, попробуйте снова.');
    });
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
			<template #body>
        <p>Вы уверены, что хотите экспортировать все данные сервиса?</p>
      </template>
      <template #footer>
        <ActionButton
          text="Нет"
          type="cancel"
          variant="flat"
          color=red
          @click="closeDialog"
        ></ActionButton>
        <ActionButton
          text="Да"
          type="confirm"
          variant="flat"
          color=green
          @click="fetchExportData"
        ></ActionButton>
      </template>
    </Dialog>
		<Dialog
      title="Импорт"
      :visible="isDialogImportVisible"
      dialogMaxWidth="30%"
      @update:visible="closeDialog"
			>
      <template #body>
        <p>Вы уверены что хотите загрузить новые данные? Это приведет к удалению всех текущих данных</p>
        <div
          class="drag-drop-area"
          @dragover.prevent
          @drop.prevent
          @drop="handleFileDrop"
        >
        <span>{{ fileName || 'Перетащите файл сюда для загрузки' }}</span>
        </div>
        <div v-if="isLoading">
          <p style="font-weight: bold">Через {{ remainingSeconds }} секунд данные в базе будут заменены данными из файла</p>          <progress :value="loadingProgress" max="100"></progress>
          <ActionButton
            text="Отменить импорт"
            type="cancel"
            variant="flat"
            color=orange
            @click="cancelImport"
          ></ActionButton>
        </div>
      </template>
      <template #footer>
        <ActionButton
          text="Нет"
          type="cancel"
          variant="flat"
          color=red
          @click="handleCancelAndReset"
        ></ActionButton>
        <ActionButton
          text="Да"
          type="confirm"
          variant="flat"
          color=green
          @click="confirmImportFile"
        ></ActionButton>
      </template>
    </Dialog>
		<Dialog
      title="Фильтры"
      :visible="isDialogFiltersVisible"
      dialogMaxWidth="30%"
      @update:visible="closeDialog"
			>
      <template #body>
        <p>Work in progress</p>
      </template>
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