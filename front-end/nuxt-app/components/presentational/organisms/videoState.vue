<template>
    <v-card
    class="video-state pa-2 d-flex align-center"
    color="orange lighten-3"
    tile
    >
        <div class="video-state-icon">
            <v-avatar
            color="white"
            size="56"
            class="mx-4 video-state-icon-avatar"
            >
                <v-icon color="black">mdi-microphone</v-icon>
            </v-avatar>

            <v-avatar
            color="white"
            size="56"
            class="mx-4 video-state-icon-avatar"
            >
                <v-icon color="black">mdi-camera</v-icon>
            </v-avatar>
        </div>

        <div class="video-state-leaving-btn">
            <Btn text="退出" color="red" />
        </div>
    </v-card>
</template>

<script>
import Btn from '~/components/presentational/atoms/btn';

export default {
    components: {
        Btn
    },
}
</script>

<style lang="scss">
.video-state {
    &-icon {
        text-align: center;
        width: 100%;
        transform: translateX(5%);

        &-avatar {
            box-shadow: 2px 4px 6px 0 rgba(0, 0, 0, .2)
        }
    }
    &-leaving-btn {
        width: 10%;
    }
}
</style>