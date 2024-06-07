<template>
  <div>
    <iframe
      :src="externalUrl"
      frameborder="0"
      :style="{ width: '100%', height: iframeHeight }"
    ></iframe>
  </div>
</template>

<script setup>
import { ref } from "vue";
const iframeHeight = "calc(100vh)";
const externalUrl =
  "https://baijiahao.baidu.com/s?id=1622874286087968883&wfr=spider&for=pc";
</script>

<style></style>