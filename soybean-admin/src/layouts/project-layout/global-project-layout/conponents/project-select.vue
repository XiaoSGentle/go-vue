<template>
  <DarkModeContainer class='p2'>
    <NSelect
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
  NThing,
  NFlex,
  SelectOption
} from 'naive-ui'
import {reactive, ref, VNode} from "vue";
import DarkModeContainer from "@/components/common/dark-mode-container.vue";

defineOptions({
  name: "ProjectSelect"
})


const options = [
  {
    label: '医学影像预览',
    value: '医学影像预览',
    description: '六字诀六字刘子六字诀六字刘子六字诀子六字诀六字刘子六字子六字诀六字刘子六字诀子六字诀六字刘子六字诀诀六字刘子六字诀六字刘子六字诀六字刘子六字诀六字刘子'
  },
  {
    label: '疼痛缓解',
    value: '疼痛缓解',
    description: '医学影像六字诀六字刘子六字诀六字刘子六字诀六字刘子六字诀六字刘子六字诀子六子六字诀六字刘子六字诀字诀六字刘子六字诀六子六字诀六字刘子六字诀子六字诀六字刘子六字诀字刘子六字诀六字刘子预览'
  },
  {
    label: '六字诀',
    value: '六字诀',
    description: '六字诀六字刘子六字诀六字刘子六字诀六字刘子六字诀六子六字诀六字刘子六字诀字子六字诀六字刘子六字诀子六字诀六字刘子六字诀刘子六字诀六字刘子六字诀六字刘子'
  },
  {
    label: '六字诀',
    value: '六字诀',
    description: '六字诀六字刘子六字诀六字刘子六字诀六字刘子六字诀六子六字诀六字刘子六字诀字子六字诀六字刘子六字诀子六字诀六字刘子六字诀刘子六字诀六字刘子六字诀六字刘子'
  },
  {
    label: '六字诀',
    value: '六字诀',
    description: '六字诀六字刘子六字诀六字刘子六字诀六字刘子六字诀六子六字诀六字刘子六字诀字子六字诀六字刘子六字诀子六字诀六字刘子六字诀刘子六字诀六字刘子六字诀六字刘子'
  }
]
const renderOptions = ({_, option}: { node: VNode; option: SelectOption }) => {
  return (
      <div class={'p2 m2 mr4 cursor-pointer transition-500  hover:border-blueGray hover:border-2 hover:rounded-lg'}>
        <NFlex>
          <NAvatar src='https://szreacool.coding.net/static/project_icon/scenery-version-2-7.svg' size={'large'}
                   class={'my1'}/>
          <NThing class={'w-80%'}>
            <NThing class={'text-17px font-bold'}>{option.label}</NThing>
            <NThing class={'text-14px line-clamp-1 text-gray'}>{option.description}</NThing>
          </NThing>
        </NFlex>
      </div>)
}

</script>


