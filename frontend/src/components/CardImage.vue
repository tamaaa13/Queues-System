<script setup></script>

<template>
  <div
    class="card bg-dark overflow-hidden max-h-[550px] md:max-h-[280px] w-full rounded-t-none relative"
  >
    <img
      class="bottom-0 w-full object-cover rounded-t-none"
      src="@/assets/images/layout/main-bg.jpeg"
      alt="Antrian"
    />
    <div class="absolute top-0 left-0 bottom-0 right-0 flex items-center justify-center p-4">
      <div class="card text-white rounded-xl mt-5">
        <div class="bg-blue-500 rounded-t-xl text-center py-4 px-7">
          <h3 class="text-white">PANGGILAN</h3>
        </div>
        <div class="rounded-b-xl py-4 px-7">
          <div class="flex flex-col justify-between items-center h-full">
            <h1 class="text-5xl font-bold">A0001</h1>
            <p class="text-muted">LOKET A</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
