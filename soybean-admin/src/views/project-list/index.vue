<script lang="tsx" setup>
import {useRouterPush} from "@/hooks/common/router";

const routerPush = useRouterPush()

const pushTarget = () => {
  routerPush.routerPush({path: "/projects/overview"})
}
</script>

<template>
  <div>
    <NButton @click="pushTarget">跳转</NButton>
  </div>
</template>

<style scoped></style>
