<template>
    <div>
        <v-app-bar elevation="4">
            <v-app-bar-nav-icon @click="drawer = true"></v-app-bar-nav-icon>
            <v-toolbar-title>test page</v-toolbar-title>
            <v-spacer></v-spacer>
            <v-btn icon>
                <v-icon>mdi-magnify</v-icon>
            </v-btn>
            <template v-slot:extension>
                <v-tabs align-with-title>
                <v-tab>Tab 1</v-tab>
                <v-tab>Tab 2</v-tab>
                <v-tab>Tab 3</v-tab>
                </v-tabs>
            </template>
        </v-app-bar>

        <v-navigation-drawer
        v-model="drawer"
        absolute
        temporary
        >
        <v-list
            nav
            dense
        >
            <v-list-item-group
            v-model="group"
            active-class="deep-purple--text text--accent-4"
            >
            <v-list-item>
                <v-list-item-icon>
                <v-icon>mdi-home</v-icon>
                </v-list-item-icon>
                <v-list-item-title>Home</v-list-item-title>
            </v-list-item>

            <v-list-item>
                <v-list-item-icon>
                <v-icon>mdi-account</v-icon>
                </v-list-item-icon>
                <v-list-item-title>Account</v-list-item-title>
            </v-list-item>
            </v-list-item-group>
        </v-list>
        </v-navigation-drawer>

        <v-carousel v-model="model">
            <v-carousel-item
            v-for="(color, i) in colors"
            :key="color"
            >
            <v-sheet
                :color="color"
                height="100%"
                tile
            >
                <v-row
                class="fill-height"
                align="center"
                justify="center"
                >
                    <div class="text-h2">
                        Slide {{ i + 1 }}
                    </div>
                </v-row>
            </v-sheet>
            </v-carousel-item>
        </v-carousel>
    </div>
</template>

<script>
export default {
    layout: "test",
    data: () => ({
      drawer: false,
      group: null,
      model: 0,
      colors: [
        'primary',
        'secondary',
        'yellow darken-2',
        'red',
        'orange',
      ],
    }),
};
</script>
