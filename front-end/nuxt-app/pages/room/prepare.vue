<template>
    <div>
        <v-text-field
        label="部屋番号"
        :rules="rules"
        hide-details="auto"
        id="room-name"
        ></v-text-field>
        <div>
            <Btn text="作成" color="orange" :clickedfn="this.joinRoom" />
            <Btn text="参加" color="blue" :clickedfn="this.joinRoom" />
        </div>
    </div>
</template>

<script>
    import Btn from '~/components/presentational/atoms/btn';

    export default {
        components: {
            Btn
        },

        methods: {
            joinRoom: function() {
                //room新規作成もjoinも同じ
                const roomName = document.querySelector('#room-name').value;
                this.$router.push('/room/' + roomName);
            }
        }
    }
</script>

<style lang="scss">

</style>
