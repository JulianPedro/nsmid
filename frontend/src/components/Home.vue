<script setup>
import {computed, onMounted, reactive} from 'vue'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { fas } from '@fortawesome/free-solid-svg-icons'

const data = reactive({
  name: "",
  utilizationGPU: 0,
  utilizationMemory: 0,
  powerDraw: 0,
  fanSpeed: 0,
  temperatureGPU: 0,
  memoryTotal: 0,
  memoryFree: 0,
  memeoryUsed: 0,
  isLoading: true,
})

const GPUImageURL = computed(() => {
    if (data.name) {
        return new URL(`../assets/images/gpus/${data.name.replaceAll(' ', '_')}.png`, import.meta.url);
    }
})

onMounted(() => {
  window.runtime.EventsOn("gpuDataUpdated", (gpuData) => {
    data.name = gpuData.static.name;
    data.utilizationGPU = gpuData.dynamic.utilization.gpu;
    data.utilizationMemory = gpuData.dynamic.utilization.memory;
    data.powerDraw = gpuData.dynamic.power.draw;
    data.fanSpeed = gpuData.dynamic.fan.speed;
    data.temperatureGPU = gpuData.dynamic.temperature.gpu;
    data.memoryTotal = gpuData.dynamic.memory.total;
    data.memoryFree = gpuData.dynamic.memory.free;
    data.memeoryUsed = gpuData.dynamic.memory.used;
    data.isLoading = false;
  });
});

</script>

<template>
  <main class="flex items-center justify-center min-h-screen">
    <div v-if="data.isLoading" class="flex flex-col items-center justify-center space-y-4">
      <div class="text-gray-500 text-3xl">
        <FontAwesomeIcon :icon="fas.faSpinner" class="animate-spin" />
      </div>
      <p class="text-gray-500">Loading...</p>
    </div>

    <div v-else class="grid gap-8">
      <p id="name" class="font-bold text-3xl text-center">{{ data.name }}</p>
      <div class="grid place-items-center">
        <img id="graphicalCard" :alt="data.name" :src="GPUImageURL" style="width: 50%;"/>
      </div>
      <div class="grid grid-cols-3 gap-8">
        <div>
          <div class="flex items-center mb-2">
            <div class="w-8 h-8 flex items-center justify-center rounded-full bg-gray-200 mr-2">
              <FontAwesomeIcon :icon="fas.faTachometerAlt" class="text-gray-700"/>
            </div>
            <p>Utilization GPU: <span id="utilizationGPU">{{ data.utilizationGPU }}%</span></p>
          </div>
          <div class="flex items-center mb-2">
            <div class="w-8 h-8 flex items-center justify-center rounded-full bg-gray-200 mr-2">
              <FontAwesomeIcon :icon="fas.faMemory" class="text-gray-700"/>
            </div>
            <p>Utilization Memory: <span id="utilizationMemory">{{ data.utilizationMemory }}%</span></p>
          </div>
        </div>
        <div>
          <div class="flex items-center mb-2">
            <div class="w-8 h-8 flex items-center justify-center rounded-full bg-gray-200 mr-2">
              <FontAwesomeIcon :icon="fas.faFan" class="text-gray-700"/>
            </div>
            <p>Fan Speed: <span id="fanSpeed">{{ data.fanSpeed }}%</span></p>
            <div v-if="!data.fanSpeed" class="flex items-center space-x-4 ml-3">
              <span class="relative inline-flex items-center justify-center w-16 h-8 text-xs font-bold text-white bg-red-800 border-2 border-gray-400 rounded-md">
                Fan Stop
              </span>
            </div>
          </div>
          <div class="flex items-center mb-2">
            <div class="w-8 h-8 flex items-center justify-center rounded-full bg-gray-200 mr-2">
              <FontAwesomeIcon :icon="fas.faTemperatureHigh" class="text-gray-700"/>
            </div>
            <p>Temperature GPU: <span id="temperatureGPU">{{ data.temperatureGPU }}°C</span></p>
          </div>
          <div class="flex items-center mb-2">
            <div class="w-8 h-8 flex items-center justify-center rounded-full bg-gray-200 mr-2">
              <FontAwesomeIcon :icon="fas.faBolt" class="text-gray-700"/>
            </div>
            <p>Power Draw: <span id="powerDraw">{{ data.powerDraw }}W</span></p>
          </div>
        </div>
        <div>
          <div class="flex items-center mb-2">
            <div class="w-8 h-8 flex items-center justify-center rounded-full bg-gray-200 mr-2">
              <FontAwesomeIcon :icon="fas.faMemory" class="text-gray-700"/>
            </div>
            <p>Memory Total: <span id="memoryTotal">{{ data.memoryTotal }}MiB</span></p>
          </div>
          <div class="flex items-center mb-2">
            <div class="w-8 h-8 flex items-center justify-center rounded-full bg-gray-200 mr-2">
              <FontAwesomeIcon :icon="fas.faMemory" class="text-gray-700"/>
            </div>
            <p>Memory Free: <span id="memoryFree">{{ data.memoryFree }}MiB</span></p>
          </div>
          <div class="flex items-center mb-2">
            <div class="w-8 h-8 flex items-center justify-center rounded-full bg-gray-200 mr-2">
              <FontAwesomeIcon :icon="fas.faMemory" class="text-gray-700"/>
            </div>
            <p>Memory Used: <span id="memoryUsed">{{ data.memeoryUsed }}MiB</span></p>
          </div>
        </div>
      </div>
    </div>    
  </main>
</template>

<style scoped>

</style>
