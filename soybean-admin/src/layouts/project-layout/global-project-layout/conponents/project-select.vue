<template>
  <DarkModeContainer class='p2'>
    <n-select
        :options="options"
        :render-option="renderOptions"
        filterable
        size="large"
    />
  </DarkModeContainer>
</template>


<script lang="tsx" setup>
import {
  NAvatar,
  NText,
  NTag,
  NThing,
  NSpace,
  SelectRenderTag,
  SelectRenderLabel, SelectGroupOption, SelectOption
} from 'naive-ui'
import {reactive, ref, VNode} from "vue";
import DarkModeContainer from "@/components/common/dark-mode-container.vue";

defineOptions({
  name: "ProjectSelect"
})


const options = [
  {
    label: '医学影像预览',
    value: '医学影像预览'
  },
  {
    label: '疼痛缓解',
    value: '疼痛缓解'
  },
  {
    label: '六字诀',
    value: '六字诀'
  }
]


const selectProject = ref("09akioni")
const renderOptions = ({node, option}: { node: VNode; option: SelectOption }) => {
  return (
      <div class={'p2'}>
        <NSpace>
          <NAvatar src='https://szreacool.coding.net/static/project_icon/scenery-version-2-7.svg'></NAvatar>
          <NThing class={'leading-relaxed'}>{option.label}</NThing>
        </NSpace>
      </div>)
}

</script>


