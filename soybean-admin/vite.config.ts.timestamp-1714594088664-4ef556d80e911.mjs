// vite.config.ts
import process3 from "node:process";
import { URL, fileURLToPath } from "node:url";
import { defineConfig, loadEnv } from "file:///E:/go-vue/soybean-admin/node_modules/.pnpm/vite@5.2.10_@types+node@20.12.7_sass@1.75.0/node_modules/vite/dist/node/index.js";
import dayjs from "file:///E:/go-vue/soybean-admin/node_modules/.pnpm/dayjs@1.11.10/node_modules/dayjs/dayjs.min.js";

// build/plugins/index.ts
import vue from "file:///E:/go-vue/soybean-admin/node_modules/.pnpm/@vitejs+plugin-vue@5.0.4_vite@5.2.10_vue@3.4.25/node_modules/@vitejs/plugin-vue/dist/index.mjs";
import vueJsx from "file:///E:/go-vue/soybean-admin/node_modules/.pnpm/@vitejs+plugin-vue-jsx@3.1.0_vite@5.2.10_vue@3.4.25/node_modules/@vitejs/plugin-vue-jsx/dist/index.mjs";
import VueDevtools from "file:///E:/go-vue/soybean-admin/node_modules/.pnpm/vite-plugin-vue-devtools@7.1.3_vite@5.2.10_vue@3.4.25/node_modules/vite-plugin-vue-devtools/dist/vite.mjs";
import progress from "file:///E:/go-vue/soybean-admin/node_modules/.pnpm/vite-plugin-progress@0.0.7_vite@5.2.10/node_modules/vite-plugin-progress/dist/index.mjs";

// build/plugins/router.ts
import ElegantVueRouter from "file:///E:/go-vue/soybean-admin/node_modules/.pnpm/@elegant-router+vue@0.3.6/node_modules/@elegant-router/vue/dist/vite.mjs";
function setupElegantRouter() {
  return ElegantVueRouter({
    layouts: {
      base: "src/layouts/base-layout/index.vue",
      blank: "src/layouts/blank-layout/index.vue"
    },
    customRoutes: {
      names: [
        "exception_403",
        "exception_404",
        "exception_500",
        "document_project",
        "document_project-link",
        "document_vue",
        "document_vite",
        "document_unocss",
        "document_naive",
        "document_antd"
      ]
    },
    routePathTransformer(routeName, routePath) {
      const key = routeName;
      if (key === "login") {
        const modules = ["pwd-login", "code-login", "register", "reset-pwd", "bind-wechat"];
        const moduleReg = modules.join("|");
        return `/login/:module(${moduleReg})?`;
      }
      return routePath;
    },
    onRouteMetaGen(routeName) {
      const key = routeName;
      const constantRoutes = ["login", "403", "404", "500"];
      const meta = {
        title: key,
        i18nKey: `route.${key}`
      };
      if (constantRoutes.includes(key)) {
        meta.constant = true;
      }
      return meta;
    }
  });
}

// build/plugins/unocss.ts
import process from "node:process";
import path from "node:path";
import unocss from "file:///E:/go-vue/soybean-admin/node_modules/.pnpm/@unocss+vite@0.59.4_vite@5.2.10/node_modules/@unocss/vite/dist/index.mjs";
import presetIcons from "file:///E:/go-vue/soybean-admin/node_modules/.pnpm/@unocss+preset-icons@0.59.4/node_modules/@unocss/preset-icons/dist/index.mjs";
import { FileSystemIconLoader } from "file:///E:/go-vue/soybean-admin/node_modules/.pnpm/@iconify+utils@2.1.23/node_modules/@iconify/utils/lib/loader/node-loaders.mjs";
function setupUnocss(viteEnv) {
  const { VITE_ICON_PREFIX, VITE_ICON_LOCAL_PREFIX } = viteEnv;
  const localIconPath = path.join(process.cwd(), "src/assets/svg-icon");
  const collectionName = VITE_ICON_LOCAL_PREFIX.replace(`${VITE_ICON_PREFIX}-`, "");
  return unocss({
    presets: [
      presetIcons({
        prefix: `${VITE_ICON_PREFIX}-`,
        scale: 1,
        extraProperties: {
          display: "inline-block"
        },
        collections: {
          [collectionName]: FileSystemIconLoader(
            localIconPath,
            (svg) => svg.replace(/^<svg\s/, '<svg width="1em" height="1em" ')
          )
        },
        warn: true
      })
    ]
  });
}

// build/plugins/unplugin.ts
import process2 from "node:process";
import path2 from "node:path";
import Icons from "file:///E:/go-vue/soybean-admin/node_modules/.pnpm/unplugin-icons@0.18.5/node_modules/unplugin-icons/dist/vite.js";
import IconsResolver from "file:///E:/go-vue/soybean-admin/node_modules/.pnpm/unplugin-icons@0.18.5/node_modules/unplugin-icons/dist/resolver.js";
import Components from "file:///E:/go-vue/soybean-admin/node_modules/.pnpm/unplugin-vue-components@0.26.0_vue@3.4.25/node_modules/unplugin-vue-components/dist/vite.js";
import { AntDesignVueResolver, NaiveUiResolver } from "file:///E:/go-vue/soybean-admin/node_modules/.pnpm/unplugin-vue-components@0.26.0_vue@3.4.25/node_modules/unplugin-vue-components/dist/resolvers.js";
import { FileSystemIconLoader as FileSystemIconLoader2 } from "file:///E:/go-vue/soybean-admin/node_modules/.pnpm/unplugin-icons@0.18.5/node_modules/unplugin-icons/dist/loaders.js";
import { createSvgIconsPlugin } from "file:///E:/go-vue/soybean-admin/node_modules/.pnpm/vite-plugin-svg-icons@2.0.1_vite@5.2.10/node_modules/vite-plugin-svg-icons/dist/index.mjs";
function setupUnplugin(viteEnv) {
  const { VITE_ICON_PREFIX, VITE_ICON_LOCAL_PREFIX } = viteEnv;
  const localIconPath = path2.join(process2.cwd(), "src/assets/svg-icon");
  const collectionName = VITE_ICON_LOCAL_PREFIX.replace(`${VITE_ICON_PREFIX}-`, "");
  const plugins = [
    Icons({
      compiler: "vue3",
      customCollections: {
        [collectionName]: FileSystemIconLoader2(
          localIconPath,
          (svg) => svg.replace(/^<svg\s/, '<svg width="1em" height="1em" ')
        )
      },
      scale: 1,
      defaultClass: "inline-block"
    }),
    Components({
      dts: "src/typings/components.d.ts",
      types: [{ from: "vue-router", names: ["RouterLink", "RouterView"] }],
      resolvers: [
        AntDesignVueResolver({
          importStyle: false
        }),
        NaiveUiResolver(),
        IconsResolver({ customCollections: [collectionName], componentPrefix: VITE_ICON_PREFIX })
      ]
    }),
    createSvgIconsPlugin({
      iconDirs: [localIconPath],
      symbolId: `${VITE_ICON_LOCAL_PREFIX}-[dir]-[name]`,
      inject: "body-last",
      customDomId: "__SVG_ICON_LOCAL__"
    })
  ];
  return plugins;
}

// build/plugins/index.ts
function setupVitePlugins(viteEnv) {
  const plugins = [
    vue({
      script: {
        defineModel: true
      }
    }),
    vueJsx(),
    VueDevtools(),
    setupElegantRouter(),
    setupUnocss(viteEnv),
    ...setupUnplugin(viteEnv),
    progress()
  ];
  return plugins;
}

// src/utils/service.ts
function createServiceConfig(env) {
  const { VITE_SERVICE_BASE_URL, VITE_OTHER_SERVICE_BASE_URL } = env;
  let other = {};
  try {
    other = JSON.parse(VITE_OTHER_SERVICE_BASE_URL);
  } catch (error) {
    console.error("VITE_OTHER_SERVICE_BASE_URL is not a valid JSON string");
  }
  const httpConfig = {
    baseURL: VITE_SERVICE_BASE_URL,
    other
  };
  const otherHttpKeys = Object.keys(httpConfig.other);
  const otherConfig = otherHttpKeys.map((key) => {
    return {
      key,
      baseURL: httpConfig.other[key],
      proxyPattern: createProxyPattern(key)
    };
  });
  const config = {
    baseURL: httpConfig.baseURL,
    proxyPattern: createProxyPattern(),
    other: otherConfig
  };
  return config;
}
function createProxyPattern(key) {
  if (!key) {
    return "/proxy-default";
  }
  return `/proxy-${key}`;
}

// build/config/proxy.ts
function createViteProxy(env, isDev) {
  const isEnableHttpProxy = isDev && env.VITE_HTTP_PROXY === "Y";
  if (!isEnableHttpProxy)
    return void 0;
  const { baseURL, proxyPattern, other } = createServiceConfig(env);
  const proxy = createProxyItem({ baseURL, proxyPattern });
  other.forEach((item) => {
    Object.assign(proxy, createProxyItem(item));
  });
  return proxy;
}
function createProxyItem(item) {
  const proxy = {};
  proxy[item.proxyPattern] = {
    target: item.baseURL,
    changeOrigin: true,
    rewrite: (path3) => path3.replace(new RegExp(`^${item.proxyPattern}`), "")
  };
  return proxy;
}

// vite.config.ts
var __vite_injected_original_import_meta_url = "file:///E:/go-vue/soybean-admin/vite.config.ts";
var vite_config_default = defineConfig((configEnv) => {
  const viteEnv = loadEnv(configEnv.mode, process3.cwd());
  const buildTime = dayjs().format("YYYY-MM-DD HH:mm:ss");
  return {
    base: viteEnv.VITE_BASE_URL,
    resolve: {
      alias: {
        "~": fileURLToPath(new URL("./", __vite_injected_original_import_meta_url)),
        "@": fileURLToPath(new URL("./src", __vite_injected_original_import_meta_url))
      }
    },
    css: {
      preprocessorOptions: {
        scss: {
          additionalData: `@use "./src/styles/scss/global.scss" as *;`
        }
      }
    },
    plugins: setupVitePlugins(viteEnv),
    define: {
      BUILD_TIME: JSON.stringify(buildTime)
    },
    server: {
      host: "0.0.0.0",
      port: 9527,
      open: true,
      proxy: createViteProxy(viteEnv, configEnv.command === "serve"),
      fs: {
        cachedChecks: false
      }
    },
    preview: {
      port: 9725
    },
    build: {
      reportCompressedSize: false,
      sourcemap: viteEnv.VITE_SOURCE_MAP === "Y",
      commonjsOptions: {
        ignoreTryCatch: false
      }
    }
  };
});
export {
  vite_config_default as default
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsidml0ZS5jb25maWcudHMiLCAiYnVpbGQvcGx1Z2lucy9pbmRleC50cyIsICJidWlsZC9wbHVnaW5zL3JvdXRlci50cyIsICJidWlsZC9wbHVnaW5zL3Vub2Nzcy50cyIsICJidWlsZC9wbHVnaW5zL3VucGx1Z2luLnRzIiwgInNyYy91dGlscy9zZXJ2aWNlLnRzIiwgImJ1aWxkL2NvbmZpZy9wcm94eS50cyJdLAogICJzb3VyY2VzQ29udGVudCI6IFsiY29uc3QgX192aXRlX2luamVjdGVkX29yaWdpbmFsX2Rpcm5hbWUgPSBcIkU6XFxcXGdvLXZ1ZVxcXFxzb3liZWFuLWFkbWluXCI7Y29uc3QgX192aXRlX2luamVjdGVkX29yaWdpbmFsX2ZpbGVuYW1lID0gXCJFOlxcXFxnby12dWVcXFxcc295YmVhbi1hZG1pblxcXFx2aXRlLmNvbmZpZy50c1wiO2NvbnN0IF9fdml0ZV9pbmplY3RlZF9vcmlnaW5hbF9pbXBvcnRfbWV0YV91cmwgPSBcImZpbGU6Ly8vRTovZ28tdnVlL3NveWJlYW4tYWRtaW4vdml0ZS5jb25maWcudHNcIjtpbXBvcnQgcHJvY2VzcyBmcm9tICdub2RlOnByb2Nlc3MnO1xuaW1wb3J0IHsgVVJMLCBmaWxlVVJMVG9QYXRoIH0gZnJvbSAnbm9kZTp1cmwnO1xuaW1wb3J0IHsgZGVmaW5lQ29uZmlnLCBsb2FkRW52IH0gZnJvbSAndml0ZSc7XG5pbXBvcnQgZGF5anMgZnJvbSAnZGF5anMnO1xuaW1wb3J0IHsgc2V0dXBWaXRlUGx1Z2lucyB9IGZyb20gJy4vYnVpbGQvcGx1Z2lucyc7XG5pbXBvcnQgeyBjcmVhdGVWaXRlUHJveHkgfSBmcm9tICcuL2J1aWxkL2NvbmZpZyc7XG5cbmV4cG9ydCBkZWZhdWx0IGRlZmluZUNvbmZpZyhjb25maWdFbnYgPT4ge1xuICBjb25zdCB2aXRlRW52ID0gbG9hZEVudihjb25maWdFbnYubW9kZSwgcHJvY2Vzcy5jd2QoKSkgYXMgdW5rbm93biBhcyBFbnYuSW1wb3J0TWV0YTtcblxuICBjb25zdCBidWlsZFRpbWUgPSBkYXlqcygpLmZvcm1hdCgnWVlZWS1NTS1ERCBISDptbTpzcycpO1xuXG4gIHJldHVybiB7XG4gICAgYmFzZTogdml0ZUVudi5WSVRFX0JBU0VfVVJMLFxuICAgIHJlc29sdmU6IHtcbiAgICAgIGFsaWFzOiB7XG4gICAgICAgICd+JzogZmlsZVVSTFRvUGF0aChuZXcgVVJMKCcuLycsIGltcG9ydC5tZXRhLnVybCkpLFxuICAgICAgICAnQCc6IGZpbGVVUkxUb1BhdGgobmV3IFVSTCgnLi9zcmMnLCBpbXBvcnQubWV0YS51cmwpKVxuICAgICAgfVxuICAgIH0sXG4gICAgY3NzOiB7XG4gICAgICBwcmVwcm9jZXNzb3JPcHRpb25zOiB7XG4gICAgICAgIHNjc3M6IHtcbiAgICAgICAgICBhZGRpdGlvbmFsRGF0YTogYEB1c2UgXCIuL3NyYy9zdHlsZXMvc2Nzcy9nbG9iYWwuc2Nzc1wiIGFzICo7YFxuICAgICAgICB9XG4gICAgICB9XG4gICAgfSxcbiAgICBwbHVnaW5zOiBzZXR1cFZpdGVQbHVnaW5zKHZpdGVFbnYpLFxuICAgIGRlZmluZToge1xuICAgICAgQlVJTERfVElNRTogSlNPTi5zdHJpbmdpZnkoYnVpbGRUaW1lKVxuICAgIH0sXG4gICAgc2VydmVyOiB7XG4gICAgICBob3N0OiAnMC4wLjAuMCcsXG4gICAgICBwb3J0OiA5NTI3LFxuICAgICAgb3BlbjogdHJ1ZSxcbiAgICAgIHByb3h5OiBjcmVhdGVWaXRlUHJveHkodml0ZUVudiwgY29uZmlnRW52LmNvbW1hbmQgPT09ICdzZXJ2ZScpLFxuICAgICAgZnM6IHtcbiAgICAgICAgY2FjaGVkQ2hlY2tzOiBmYWxzZVxuICAgICAgfVxuICAgIH0sXG4gICAgcHJldmlldzoge1xuICAgICAgcG9ydDogOTcyNVxuICAgIH0sXG4gICAgYnVpbGQ6IHtcbiAgICAgIHJlcG9ydENvbXByZXNzZWRTaXplOiBmYWxzZSxcbiAgICAgIHNvdXJjZW1hcDogdml0ZUVudi5WSVRFX1NPVVJDRV9NQVAgPT09ICdZJyxcbiAgICAgIGNvbW1vbmpzT3B0aW9uczoge1xuICAgICAgICBpZ25vcmVUcnlDYXRjaDogZmFsc2VcbiAgICAgIH1cbiAgICB9XG4gIH07XG59KTtcbiIsICJjb25zdCBfX3ZpdGVfaW5qZWN0ZWRfb3JpZ2luYWxfZGlybmFtZSA9IFwiRTpcXFxcZ28tdnVlXFxcXHNveWJlYW4tYWRtaW5cXFxcYnVpbGRcXFxccGx1Z2luc1wiO2NvbnN0IF9fdml0ZV9pbmplY3RlZF9vcmlnaW5hbF9maWxlbmFtZSA9IFwiRTpcXFxcZ28tdnVlXFxcXHNveWJlYW4tYWRtaW5cXFxcYnVpbGRcXFxccGx1Z2luc1xcXFxpbmRleC50c1wiO2NvbnN0IF9fdml0ZV9pbmplY3RlZF9vcmlnaW5hbF9pbXBvcnRfbWV0YV91cmwgPSBcImZpbGU6Ly8vRTovZ28tdnVlL3NveWJlYW4tYWRtaW4vYnVpbGQvcGx1Z2lucy9pbmRleC50c1wiO2ltcG9ydCB0eXBlIHsgUGx1Z2luT3B0aW9uIH0gZnJvbSAndml0ZSc7XG5pbXBvcnQgdnVlIGZyb20gJ0B2aXRlanMvcGx1Z2luLXZ1ZSc7XG5pbXBvcnQgdnVlSnN4IGZyb20gJ0B2aXRlanMvcGx1Z2luLXZ1ZS1qc3gnO1xuaW1wb3J0IFZ1ZURldnRvb2xzIGZyb20gJ3ZpdGUtcGx1Z2luLXZ1ZS1kZXZ0b29scyc7XG5pbXBvcnQgcHJvZ3Jlc3MgZnJvbSAndml0ZS1wbHVnaW4tcHJvZ3Jlc3MnO1xuaW1wb3J0IHsgc2V0dXBFbGVnYW50Um91dGVyIH0gZnJvbSAnLi9yb3V0ZXInO1xuaW1wb3J0IHsgc2V0dXBVbm9jc3MgfSBmcm9tICcuL3Vub2Nzcyc7XG5pbXBvcnQgeyBzZXR1cFVucGx1Z2luIH0gZnJvbSAnLi91bnBsdWdpbic7XG5cbmV4cG9ydCBmdW5jdGlvbiBzZXR1cFZpdGVQbHVnaW5zKHZpdGVFbnY6IEVudi5JbXBvcnRNZXRhKSB7XG4gIGNvbnN0IHBsdWdpbnM6IFBsdWdpbk9wdGlvbiA9IFtcbiAgICB2dWUoe1xuICAgICAgc2NyaXB0OiB7XG4gICAgICAgIGRlZmluZU1vZGVsOiB0cnVlXG4gICAgICB9XG4gICAgfSksXG4gICAgdnVlSnN4KCksXG4gICAgVnVlRGV2dG9vbHMoKSxcbiAgICBzZXR1cEVsZWdhbnRSb3V0ZXIoKSxcbiAgICBzZXR1cFVub2Nzcyh2aXRlRW52KSxcbiAgICAuLi5zZXR1cFVucGx1Z2luKHZpdGVFbnYpLFxuICAgIHByb2dyZXNzKClcbiAgXTtcblxuICByZXR1cm4gcGx1Z2lucztcbn1cbiIsICJjb25zdCBfX3ZpdGVfaW5qZWN0ZWRfb3JpZ2luYWxfZGlybmFtZSA9IFwiRTpcXFxcZ28tdnVlXFxcXHNveWJlYW4tYWRtaW5cXFxcYnVpbGRcXFxccGx1Z2luc1wiO2NvbnN0IF9fdml0ZV9pbmplY3RlZF9vcmlnaW5hbF9maWxlbmFtZSA9IFwiRTpcXFxcZ28tdnVlXFxcXHNveWJlYW4tYWRtaW5cXFxcYnVpbGRcXFxccGx1Z2luc1xcXFxyb3V0ZXIudHNcIjtjb25zdCBfX3ZpdGVfaW5qZWN0ZWRfb3JpZ2luYWxfaW1wb3J0X21ldGFfdXJsID0gXCJmaWxlOi8vL0U6L2dvLXZ1ZS9zb3liZWFuLWFkbWluL2J1aWxkL3BsdWdpbnMvcm91dGVyLnRzXCI7aW1wb3J0IHR5cGUgeyBSb3V0ZU1ldGEgfSBmcm9tICd2dWUtcm91dGVyJztcbmltcG9ydCBFbGVnYW50VnVlUm91dGVyIGZyb20gJ0BlbGVnYW50LXJvdXRlci92dWUvdml0ZSc7XG5pbXBvcnQgdHlwZSB7IFJvdXRlS2V5IH0gZnJvbSAnQGVsZWdhbnQtcm91dGVyL3R5cGVzJztcblxuZXhwb3J0IGZ1bmN0aW9uIHNldHVwRWxlZ2FudFJvdXRlcigpIHtcbiAgcmV0dXJuIEVsZWdhbnRWdWVSb3V0ZXIoe1xuICAgIGxheW91dHM6IHtcbiAgICAgIGJhc2U6ICdzcmMvbGF5b3V0cy9iYXNlLWxheW91dC9pbmRleC52dWUnLFxuICAgICAgYmxhbms6ICdzcmMvbGF5b3V0cy9ibGFuay1sYXlvdXQvaW5kZXgudnVlJ1xuICAgIH0sXG4gICAgY3VzdG9tUm91dGVzOiB7XG4gICAgICBuYW1lczogW1xuICAgICAgICAnZXhjZXB0aW9uXzQwMycsXG4gICAgICAgICdleGNlcHRpb25fNDA0JyxcbiAgICAgICAgJ2V4Y2VwdGlvbl81MDAnLFxuICAgICAgICAnZG9jdW1lbnRfcHJvamVjdCcsXG4gICAgICAgICdkb2N1bWVudF9wcm9qZWN0LWxpbmsnLFxuICAgICAgICAnZG9jdW1lbnRfdnVlJyxcbiAgICAgICAgJ2RvY3VtZW50X3ZpdGUnLFxuICAgICAgICAnZG9jdW1lbnRfdW5vY3NzJyxcbiAgICAgICAgJ2RvY3VtZW50X25haXZlJyxcbiAgICAgICAgJ2RvY3VtZW50X2FudGQnXG4gICAgICBdXG4gICAgfSxcbiAgICByb3V0ZVBhdGhUcmFuc2Zvcm1lcihyb3V0ZU5hbWUsIHJvdXRlUGF0aCkge1xuICAgICAgY29uc3Qga2V5ID0gcm91dGVOYW1lIGFzIFJvdXRlS2V5O1xuXG4gICAgICBpZiAoa2V5ID09PSAnbG9naW4nKSB7XG4gICAgICAgIGNvbnN0IG1vZHVsZXM6IFVuaW9uS2V5LkxvZ2luTW9kdWxlW10gPSBbJ3B3ZC1sb2dpbicsICdjb2RlLWxvZ2luJywgJ3JlZ2lzdGVyJywgJ3Jlc2V0LXB3ZCcsICdiaW5kLXdlY2hhdCddO1xuXG4gICAgICAgIGNvbnN0IG1vZHVsZVJlZyA9IG1vZHVsZXMuam9pbignfCcpO1xuXG4gICAgICAgIHJldHVybiBgL2xvZ2luLzptb2R1bGUoJHttb2R1bGVSZWd9KT9gO1xuICAgICAgfVxuXG4gICAgICByZXR1cm4gcm91dGVQYXRoO1xuICAgIH0sXG4gICAgb25Sb3V0ZU1ldGFHZW4ocm91dGVOYW1lKSB7XG4gICAgICBjb25zdCBrZXkgPSByb3V0ZU5hbWUgYXMgUm91dGVLZXk7XG5cbiAgICAgIGNvbnN0IGNvbnN0YW50Um91dGVzOiBSb3V0ZUtleVtdID0gWydsb2dpbicsICc0MDMnLCAnNDA0JywgJzUwMCddO1xuXG4gICAgICBjb25zdCBtZXRhOiBQYXJ0aWFsPFJvdXRlTWV0YT4gPSB7XG4gICAgICAgIHRpdGxlOiBrZXksXG4gICAgICAgIGkxOG5LZXk6IGByb3V0ZS4ke2tleX1gIGFzIEFwcC5JMThuLkkxOG5LZXlcbiAgICAgIH07XG5cbiAgICAgIGlmIChjb25zdGFudFJvdXRlcy5pbmNsdWRlcyhrZXkpKSB7XG4gICAgICAgIG1ldGEuY29uc3RhbnQgPSB0cnVlO1xuICAgICAgfVxuXG4gICAgICByZXR1cm4gbWV0YTtcbiAgICB9XG4gIH0pO1xufVxuIiwgImNvbnN0IF9fdml0ZV9pbmplY3RlZF9vcmlnaW5hbF9kaXJuYW1lID0gXCJFOlxcXFxnby12dWVcXFxcc295YmVhbi1hZG1pblxcXFxidWlsZFxcXFxwbHVnaW5zXCI7Y29uc3QgX192aXRlX2luamVjdGVkX29yaWdpbmFsX2ZpbGVuYW1lID0gXCJFOlxcXFxnby12dWVcXFxcc295YmVhbi1hZG1pblxcXFxidWlsZFxcXFxwbHVnaW5zXFxcXHVub2Nzcy50c1wiO2NvbnN0IF9fdml0ZV9pbmplY3RlZF9vcmlnaW5hbF9pbXBvcnRfbWV0YV91cmwgPSBcImZpbGU6Ly8vRTovZ28tdnVlL3NveWJlYW4tYWRtaW4vYnVpbGQvcGx1Z2lucy91bm9jc3MudHNcIjtpbXBvcnQgcHJvY2VzcyBmcm9tICdub2RlOnByb2Nlc3MnO1xuaW1wb3J0IHBhdGggZnJvbSAnbm9kZTpwYXRoJztcbmltcG9ydCB1bm9jc3MgZnJvbSAnQHVub2Nzcy92aXRlJztcbmltcG9ydCBwcmVzZXRJY29ucyBmcm9tICdAdW5vY3NzL3ByZXNldC1pY29ucyc7XG5pbXBvcnQgeyBGaWxlU3lzdGVtSWNvbkxvYWRlciB9IGZyb20gJ0BpY29uaWZ5L3V0aWxzL2xpYi9sb2FkZXIvbm9kZS1sb2FkZXJzJztcblxuZXhwb3J0IGZ1bmN0aW9uIHNldHVwVW5vY3NzKHZpdGVFbnY6IEVudi5JbXBvcnRNZXRhKSB7XG4gIGNvbnN0IHsgVklURV9JQ09OX1BSRUZJWCwgVklURV9JQ09OX0xPQ0FMX1BSRUZJWCB9ID0gdml0ZUVudjtcblxuICBjb25zdCBsb2NhbEljb25QYXRoID0gcGF0aC5qb2luKHByb2Nlc3MuY3dkKCksICdzcmMvYXNzZXRzL3N2Zy1pY29uJyk7XG5cbiAgLyoqIFRoZSBuYW1lIG9mIHRoZSBsb2NhbCBpY29uIGNvbGxlY3Rpb24gKi9cbiAgY29uc3QgY29sbGVjdGlvbk5hbWUgPSBWSVRFX0lDT05fTE9DQUxfUFJFRklYLnJlcGxhY2UoYCR7VklURV9JQ09OX1BSRUZJWH0tYCwgJycpO1xuXG4gIHJldHVybiB1bm9jc3Moe1xuICAgIHByZXNldHM6IFtcbiAgICAgIHByZXNldEljb25zKHtcbiAgICAgICAgcHJlZml4OiBgJHtWSVRFX0lDT05fUFJFRklYfS1gLFxuICAgICAgICBzY2FsZTogMSxcbiAgICAgICAgZXh0cmFQcm9wZXJ0aWVzOiB7XG4gICAgICAgICAgZGlzcGxheTogJ2lubGluZS1ibG9jaydcbiAgICAgICAgfSxcbiAgICAgICAgY29sbGVjdGlvbnM6IHtcbiAgICAgICAgICBbY29sbGVjdGlvbk5hbWVdOiBGaWxlU3lzdGVtSWNvbkxvYWRlcihsb2NhbEljb25QYXRoLCBzdmcgPT5cbiAgICAgICAgICAgIHN2Zy5yZXBsYWNlKC9ePHN2Z1xccy8sICc8c3ZnIHdpZHRoPVwiMWVtXCIgaGVpZ2h0PVwiMWVtXCIgJylcbiAgICAgICAgICApXG4gICAgICAgIH0sXG4gICAgICAgIHdhcm46IHRydWVcbiAgICAgIH0pXG4gICAgXVxuICB9KTtcbn1cbiIsICJjb25zdCBfX3ZpdGVfaW5qZWN0ZWRfb3JpZ2luYWxfZGlybmFtZSA9IFwiRTpcXFxcZ28tdnVlXFxcXHNveWJlYW4tYWRtaW5cXFxcYnVpbGRcXFxccGx1Z2luc1wiO2NvbnN0IF9fdml0ZV9pbmplY3RlZF9vcmlnaW5hbF9maWxlbmFtZSA9IFwiRTpcXFxcZ28tdnVlXFxcXHNveWJlYW4tYWRtaW5cXFxcYnVpbGRcXFxccGx1Z2luc1xcXFx1bnBsdWdpbi50c1wiO2NvbnN0IF9fdml0ZV9pbmplY3RlZF9vcmlnaW5hbF9pbXBvcnRfbWV0YV91cmwgPSBcImZpbGU6Ly8vRTovZ28tdnVlL3NveWJlYW4tYWRtaW4vYnVpbGQvcGx1Z2lucy91bnBsdWdpbi50c1wiO2ltcG9ydCBwcm9jZXNzIGZyb20gJ25vZGU6cHJvY2Vzcyc7XG5pbXBvcnQgcGF0aCBmcm9tICdub2RlOnBhdGgnO1xuaW1wb3J0IHR5cGUgeyBQbHVnaW5PcHRpb24gfSBmcm9tICd2aXRlJztcbmltcG9ydCBJY29ucyBmcm9tICd1bnBsdWdpbi1pY29ucy92aXRlJztcbmltcG9ydCBJY29uc1Jlc29sdmVyIGZyb20gJ3VucGx1Z2luLWljb25zL3Jlc29sdmVyJztcbmltcG9ydCBDb21wb25lbnRzIGZyb20gJ3VucGx1Z2luLXZ1ZS1jb21wb25lbnRzL3ZpdGUnO1xuaW1wb3J0IHsgQW50RGVzaWduVnVlUmVzb2x2ZXIsIE5haXZlVWlSZXNvbHZlciB9IGZyb20gJ3VucGx1Z2luLXZ1ZS1jb21wb25lbnRzL3Jlc29sdmVycyc7XG5pbXBvcnQgeyBGaWxlU3lzdGVtSWNvbkxvYWRlciB9IGZyb20gJ3VucGx1Z2luLWljb25zL2xvYWRlcnMnO1xuaW1wb3J0IHsgY3JlYXRlU3ZnSWNvbnNQbHVnaW4gfSBmcm9tICd2aXRlLXBsdWdpbi1zdmctaWNvbnMnO1xuXG5leHBvcnQgZnVuY3Rpb24gc2V0dXBVbnBsdWdpbih2aXRlRW52OiBFbnYuSW1wb3J0TWV0YSkge1xuICBjb25zdCB7IFZJVEVfSUNPTl9QUkVGSVgsIFZJVEVfSUNPTl9MT0NBTF9QUkVGSVggfSA9IHZpdGVFbnY7XG5cbiAgY29uc3QgbG9jYWxJY29uUGF0aCA9IHBhdGguam9pbihwcm9jZXNzLmN3ZCgpLCAnc3JjL2Fzc2V0cy9zdmctaWNvbicpO1xuXG4gIC8qKiBUaGUgbmFtZSBvZiB0aGUgbG9jYWwgaWNvbiBjb2xsZWN0aW9uICovXG4gIGNvbnN0IGNvbGxlY3Rpb25OYW1lID0gVklURV9JQ09OX0xPQ0FMX1BSRUZJWC5yZXBsYWNlKGAke1ZJVEVfSUNPTl9QUkVGSVh9LWAsICcnKTtcblxuICBjb25zdCBwbHVnaW5zOiBQbHVnaW5PcHRpb25bXSA9IFtcbiAgICBJY29ucyh7XG4gICAgICBjb21waWxlcjogJ3Z1ZTMnLFxuICAgICAgY3VzdG9tQ29sbGVjdGlvbnM6IHtcbiAgICAgICAgW2NvbGxlY3Rpb25OYW1lXTogRmlsZVN5c3RlbUljb25Mb2FkZXIobG9jYWxJY29uUGF0aCwgc3ZnID0+XG4gICAgICAgICAgc3ZnLnJlcGxhY2UoL148c3ZnXFxzLywgJzxzdmcgd2lkdGg9XCIxZW1cIiBoZWlnaHQ9XCIxZW1cIiAnKVxuICAgICAgICApXG4gICAgICB9LFxuICAgICAgc2NhbGU6IDEsXG4gICAgICBkZWZhdWx0Q2xhc3M6ICdpbmxpbmUtYmxvY2snXG4gICAgfSksXG4gICAgQ29tcG9uZW50cyh7XG4gICAgICBkdHM6ICdzcmMvdHlwaW5ncy9jb21wb25lbnRzLmQudHMnLFxuICAgICAgdHlwZXM6IFt7IGZyb206ICd2dWUtcm91dGVyJywgbmFtZXM6IFsnUm91dGVyTGluaycsICdSb3V0ZXJWaWV3J10gfV0sXG4gICAgICByZXNvbHZlcnM6IFtcbiAgICAgICAgQW50RGVzaWduVnVlUmVzb2x2ZXIoe1xuICAgICAgICAgIGltcG9ydFN0eWxlOiBmYWxzZVxuICAgICAgICB9KSxcbiAgICAgICAgTmFpdmVVaVJlc29sdmVyKCksXG4gICAgICAgIEljb25zUmVzb2x2ZXIoeyBjdXN0b21Db2xsZWN0aW9uczogW2NvbGxlY3Rpb25OYW1lXSwgY29tcG9uZW50UHJlZml4OiBWSVRFX0lDT05fUFJFRklYIH0pXG4gICAgICBdXG4gICAgfSksXG4gICAgY3JlYXRlU3ZnSWNvbnNQbHVnaW4oe1xuICAgICAgaWNvbkRpcnM6IFtsb2NhbEljb25QYXRoXSxcbiAgICAgIHN5bWJvbElkOiBgJHtWSVRFX0lDT05fTE9DQUxfUFJFRklYfS1bZGlyXS1bbmFtZV1gLFxuICAgICAgaW5qZWN0OiAnYm9keS1sYXN0JyxcbiAgICAgIGN1c3RvbURvbUlkOiAnX19TVkdfSUNPTl9MT0NBTF9fJ1xuICAgIH0pXG4gIF07XG5cbiAgcmV0dXJuIHBsdWdpbnM7XG59XG4iLCAiY29uc3QgX192aXRlX2luamVjdGVkX29yaWdpbmFsX2Rpcm5hbWUgPSBcIkU6XFxcXGdvLXZ1ZVxcXFxzb3liZWFuLWFkbWluXFxcXHNyY1xcXFx1dGlsc1wiO2NvbnN0IF9fdml0ZV9pbmplY3RlZF9vcmlnaW5hbF9maWxlbmFtZSA9IFwiRTpcXFxcZ28tdnVlXFxcXHNveWJlYW4tYWRtaW5cXFxcc3JjXFxcXHV0aWxzXFxcXHNlcnZpY2UudHNcIjtjb25zdCBfX3ZpdGVfaW5qZWN0ZWRfb3JpZ2luYWxfaW1wb3J0X21ldGFfdXJsID0gXCJmaWxlOi8vL0U6L2dvLXZ1ZS9zb3liZWFuLWFkbWluL3NyYy91dGlscy9zZXJ2aWNlLnRzXCI7LyoqXG4gKiBDcmVhdGUgc2VydmljZSBjb25maWcgYnkgY3VycmVudCBlbnZcbiAqXG4gKiBAcGFyYW0gZW52IFRoZSBjdXJyZW50IGVudlxuICovXG5leHBvcnQgZnVuY3Rpb24gY3JlYXRlU2VydmljZUNvbmZpZyhlbnY6IEVudi5JbXBvcnRNZXRhKSB7XG4gIGNvbnN0IHsgVklURV9TRVJWSUNFX0JBU0VfVVJMLCBWSVRFX09USEVSX1NFUlZJQ0VfQkFTRV9VUkwgfSA9IGVudjtcblxuICBsZXQgb3RoZXIgPSB7fSBhcyBSZWNvcmQ8QXBwLlNlcnZpY2UuT3RoZXJCYXNlVVJMS2V5LCBzdHJpbmc+O1xuICB0cnkge1xuICAgIG90aGVyID0gSlNPTi5wYXJzZShWSVRFX09USEVSX1NFUlZJQ0VfQkFTRV9VUkwpO1xuICB9IGNhdGNoIChlcnJvcikge1xuICAgIC8vIGVzbGludC1kaXNhYmxlLW5leHQtbGluZSBuby1jb25zb2xlXG4gICAgY29uc29sZS5lcnJvcignVklURV9PVEhFUl9TRVJWSUNFX0JBU0VfVVJMIGlzIG5vdCBhIHZhbGlkIEpTT04gc3RyaW5nJyk7XG4gIH1cblxuICBjb25zdCBodHRwQ29uZmlnOiBBcHAuU2VydmljZS5TaW1wbGVTZXJ2aWNlQ29uZmlnID0ge1xuICAgIGJhc2VVUkw6IFZJVEVfU0VSVklDRV9CQVNFX1VSTCxcbiAgICBvdGhlclxuICB9O1xuXG4gIGNvbnN0IG90aGVySHR0cEtleXMgPSBPYmplY3Qua2V5cyhodHRwQ29uZmlnLm90aGVyKSBhcyBBcHAuU2VydmljZS5PdGhlckJhc2VVUkxLZXlbXTtcblxuICBjb25zdCBvdGhlckNvbmZpZzogQXBwLlNlcnZpY2UuT3RoZXJTZXJ2aWNlQ29uZmlnSXRlbVtdID0gb3RoZXJIdHRwS2V5cy5tYXAoa2V5ID0+IHtcbiAgICByZXR1cm4ge1xuICAgICAga2V5LFxuICAgICAgYmFzZVVSTDogaHR0cENvbmZpZy5vdGhlcltrZXldLFxuICAgICAgcHJveHlQYXR0ZXJuOiBjcmVhdGVQcm94eVBhdHRlcm4oa2V5KVxuICAgIH07XG4gIH0pO1xuXG4gIGNvbnN0IGNvbmZpZzogQXBwLlNlcnZpY2UuU2VydmljZUNvbmZpZyA9IHtcbiAgICBiYXNlVVJMOiBodHRwQ29uZmlnLmJhc2VVUkwsXG4gICAgcHJveHlQYXR0ZXJuOiBjcmVhdGVQcm94eVBhdHRlcm4oKSxcbiAgICBvdGhlcjogb3RoZXJDb25maWdcbiAgfTtcblxuICByZXR1cm4gY29uZmlnO1xufVxuXG4vKipcbiAqIGdldCBiYWNrZW5kIHNlcnZpY2UgYmFzZSB1cmxcbiAqXG4gKiBAcGFyYW0gZW52IC0gdGhlIGN1cnJlbnQgZW52XG4gKiBAcGFyYW0gaXNQcm94eSAtIGlmIHVzZSBwcm94eVxuICovXG5leHBvcnQgZnVuY3Rpb24gZ2V0U2VydmljZUJhc2VVUkwoZW52OiBFbnYuSW1wb3J0TWV0YSwgaXNQcm94eTogYm9vbGVhbikge1xuICBjb25zdCB7IGJhc2VVUkwsIG90aGVyIH0gPSBjcmVhdGVTZXJ2aWNlQ29uZmlnKGVudik7XG5cbiAgY29uc3Qgb3RoZXJCYXNlVVJMID0ge30gYXMgUmVjb3JkPEFwcC5TZXJ2aWNlLk90aGVyQmFzZVVSTEtleSwgc3RyaW5nPjtcblxuICBvdGhlci5mb3JFYWNoKGl0ZW0gPT4ge1xuICAgIG90aGVyQmFzZVVSTFtpdGVtLmtleV0gPSBpc1Byb3h5ID8gaXRlbS5wcm94eVBhdHRlcm4gOiBpdGVtLmJhc2VVUkw7XG4gIH0pO1xuXG4gIHJldHVybiB7XG4gICAgYmFzZVVSTDogaXNQcm94eSA/IGNyZWF0ZVByb3h5UGF0dGVybigpIDogYmFzZVVSTCxcbiAgICBvdGhlckJhc2VVUkxcbiAgfTtcbn1cblxuLyoqXG4gKiBHZXQgcHJveHkgcGF0dGVybiBvZiBiYWNrZW5kIHNlcnZpY2UgYmFzZSB1cmxcbiAqXG4gKiBAcGFyYW0ga2V5IElmIG5vdCBzZXQsIHdpbGwgdXNlIHRoZSBkZWZhdWx0IGtleVxuICovXG5mdW5jdGlvbiBjcmVhdGVQcm94eVBhdHRlcm4oa2V5PzogQXBwLlNlcnZpY2UuT3RoZXJCYXNlVVJMS2V5KSB7XG4gIGlmICgha2V5KSB7XG4gICAgcmV0dXJuICcvcHJveHktZGVmYXVsdCc7XG4gIH1cblxuICByZXR1cm4gYC9wcm94eS0ke2tleX1gO1xufVxuIiwgImNvbnN0IF9fdml0ZV9pbmplY3RlZF9vcmlnaW5hbF9kaXJuYW1lID0gXCJFOlxcXFxnby12dWVcXFxcc295YmVhbi1hZG1pblxcXFxidWlsZFxcXFxjb25maWdcIjtjb25zdCBfX3ZpdGVfaW5qZWN0ZWRfb3JpZ2luYWxfZmlsZW5hbWUgPSBcIkU6XFxcXGdvLXZ1ZVxcXFxzb3liZWFuLWFkbWluXFxcXGJ1aWxkXFxcXGNvbmZpZ1xcXFxwcm94eS50c1wiO2NvbnN0IF9fdml0ZV9pbmplY3RlZF9vcmlnaW5hbF9pbXBvcnRfbWV0YV91cmwgPSBcImZpbGU6Ly8vRTovZ28tdnVlL3NveWJlYW4tYWRtaW4vYnVpbGQvY29uZmlnL3Byb3h5LnRzXCI7aW1wb3J0IHR5cGUgeyBQcm94eU9wdGlvbnMgfSBmcm9tICd2aXRlJztcbmltcG9ydCB7IGNyZWF0ZVNlcnZpY2VDb25maWcgfSBmcm9tICcuLi8uLi9zcmMvdXRpbHMvc2VydmljZSc7XG5cbi8qKlxuICogU2V0IGh0dHAgcHJveHlcbiAqXG4gKiBAcGFyYW0gZW52IC0gVGhlIGN1cnJlbnQgZW52XG4gKiBAcGFyYW0gaXNEZXYgLSBJcyBkZXZlbG9wbWVudCBlbnZpcm9ubWVudFxuICovXG5leHBvcnQgZnVuY3Rpb24gY3JlYXRlVml0ZVByb3h5KGVudjogRW52LkltcG9ydE1ldGEsIGlzRGV2OiBib29sZWFuKSB7XG4gIGNvbnN0IGlzRW5hYmxlSHR0cFByb3h5ID0gaXNEZXYgJiYgZW52LlZJVEVfSFRUUF9QUk9YWSA9PT0gJ1knO1xuXG4gIGlmICghaXNFbmFibGVIdHRwUHJveHkpIHJldHVybiB1bmRlZmluZWQ7XG5cbiAgY29uc3QgeyBiYXNlVVJMLCBwcm94eVBhdHRlcm4sIG90aGVyIH0gPSBjcmVhdGVTZXJ2aWNlQ29uZmlnKGVudik7XG5cbiAgY29uc3QgcHJveHk6IFJlY29yZDxzdHJpbmcsIFByb3h5T3B0aW9ucz4gPSBjcmVhdGVQcm94eUl0ZW0oeyBiYXNlVVJMLCBwcm94eVBhdHRlcm4gfSk7XG5cbiAgb3RoZXIuZm9yRWFjaChpdGVtID0+IHtcbiAgICBPYmplY3QuYXNzaWduKHByb3h5LCBjcmVhdGVQcm94eUl0ZW0oaXRlbSkpO1xuICB9KTtcblxuICByZXR1cm4gcHJveHk7XG59XG5cbmZ1bmN0aW9uIGNyZWF0ZVByb3h5SXRlbShpdGVtOiBBcHAuU2VydmljZS5TZXJ2aWNlQ29uZmlnSXRlbSkge1xuICBjb25zdCBwcm94eTogUmVjb3JkPHN0cmluZywgUHJveHlPcHRpb25zPiA9IHt9O1xuXG4gIHByb3h5W2l0ZW0ucHJveHlQYXR0ZXJuXSA9IHtcbiAgICB0YXJnZXQ6IGl0ZW0uYmFzZVVSTCxcbiAgICBjaGFuZ2VPcmlnaW46IHRydWUsXG4gICAgcmV3cml0ZTogcGF0aCA9PiBwYXRoLnJlcGxhY2UobmV3IFJlZ0V4cChgXiR7aXRlbS5wcm94eVBhdHRlcm59YCksICcnKVxuICB9O1xuXG4gIHJldHVybiBwcm94eTtcbn1cbiJdLAogICJtYXBwaW5ncyI6ICI7QUFBNlAsT0FBT0EsY0FBYTtBQUNqUixTQUFTLEtBQUsscUJBQXFCO0FBQ25DLFNBQVMsY0FBYyxlQUFlO0FBQ3RDLE9BQU8sV0FBVzs7O0FDRmxCLE9BQU8sU0FBUztBQUNo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QUFDdkMsUUFBTSxVQUFVLFFBQVEsVUFBVSxNQUFNQyxTQUFRLElBQUksQ0FBQztBQUVyRCxRQUFNLFlBQVksTUFBTSxFQUFFLE9BQU8scUJBQXFCO0FBRXRELFNBQU87QUFBQSxJQUNMLE1BQU0sUUFBUTtBQUFBLElBQ2QsU0FBUztBQUFBLE1BQ1AsT0FBTztBQUFBLFFBQ0wsS0FBSyxjQUFjLElBQUksSUFBSSxNQUFNLHdDQUFlLENBQUM7QUFBQSxRQUNqRCxLQUFLLGNBQWMsSUFBSSxJQUFJLFNBQVMsd0NBQWUsQ0FBQztBQUFBLE1BQ3REO0FBQUEsSUFDRjtBQUFBLElBQ0EsS0FBSztBQUFBLE1BQ0gscUJBQXFCO0FBQUEsUUFDbkIsTUFBTTtBQUFBLFVBQ0osZ0JBQWdCO0FBQUEsUUFDbEI7QUFBQSxNQUNGO0FBQUEsSUFDRjtBQUFBLElBQ0EsU0FBUyxpQkFBaUIsT0FBTztBQUFBLElBQ2pDLFFBQVE7QUFBQSxNQUNOLFlBQVksS0FBSyxVQUFVLFNBQVM7QUFBQSxJQUN0QztBQUFBLElBQ0EsUUFBUTtBQUFBLE1BQ04sTUFBTTtBQUFBLE1BQ04sTUFBTTtBQUFBLE1BQ04sTUFBTTtBQUFBLE1BQ04sT0FBTyxnQkFBZ0IsU0FBUyxVQUFVLFlBQVksT0FBTztBQUFBLE1BQzdELElBQUk7QUFBQSxRQUNGLGNBQWM7QUFBQSxNQUNoQjtBQUFBLElBQ0Y7QUFBQSxJQUNBLFNBQVM7QUFBQSxNQUNQLE1BQU07QUFBQSxJQUNSO0FBQUEsSUFDQSxPQUFPO0FBQUEsTUFDTCxzQkFBc0I7QUFBQSxNQUN0QixXQUFXLFFBQVEsb0JBQW9CO0FBQUEsTUFDdkMsaUJBQWlCO0FBQUEsUUFDZixnQkFBZ0I7QUFBQSxNQUNsQjtBQUFBLElBQ0Y7QUFBQSxFQUNGO0FBQ0YsQ0FBQzsiLAogICJuYW1lcyI6IFsicHJvY2VzcyIsICJwcm9jZXNzIiwgInBhdGgiLCAiRmlsZVN5c3RlbUljb25Mb2FkZXIiLCAicGF0aCIsICJwcm9jZXNzIiwgIkZpbGVTeXN0ZW1JY29uTG9hZGVyIiwgInBhdGgiLCAicHJvY2VzcyJdCn0K
