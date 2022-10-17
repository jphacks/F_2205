<template>
    <section>
        <section id="video-wrap"> 
            <div>
                <video id="my-video" width="400px" autoplay muted playsinline></video>
            </div>
        </section>
        
        <Btn text="接続" color="orange" :clickedfn="this.roomConnection" />

        <!-- ビデオステータスバー -->
        <VideoState :leavingFn="this.roomLeaving" />
        <!-- ビデオステータスバー -->
    </section>
</template>

<script>
import Peer from 'skyway-js'

import VideoState from '~/components/presentational/organisms/videoState';
import Btn from '~/components/presentational/atoms/btn';

export default {
    components: {
        VideoState,
        Btn
    },

    data() {
        return {
            APIKey: '5152bad7-4798-40b1-986a-a7e8f164a8a3',
            localStream: null,
            peer: null
        }
    },

    methods: {
        setSkywayEventListener: function(mediaConnection) {
            mediaConnection.on('stream', stream => {
                // video要素にカメラ映像をセットして再生
                this.addVideo(stream)
                const videoElm = document.getElementById(stream.id)
                videoElm.srcObject = stream;
                // videoElm.play();
            });

            //ルームメンバーが退出したときに発火
            mediaConnection.on('peerLeave', peerId => {
                this.removeVideo(peerId);
            });

            //何らかのエラーが発生した場合に発火
            this.peer.on('error', err => {
                alert(err.message);
            });
        },

        roomConnection: function() {
            //ルーム作成 or ルーム参加
            console.log(this.peer);
            const roomName = this.$route.params.id;
            const mediaConnection = this.peer.joinRoom(roomName, {mode: 'sfu', stream: this.localStream});
            this.setSkywayEventListener(mediaConnection);
        },

        roomLeaving: function() {
            //ルーム退出
            this.peer.destroy();
            alert('退出しました');
            this.$router.push('/room/prepare');
        },

        addVideo: function(stream) {
            const videoDom = document.createElement('video');
            videoDom.setAttribute('id', stream.peerId);
            videoDom.srcObject = stream;
            videoDom.play();
            document.getElementById('video-wrap').append(videoDom);
        },
        removeVideo: function(peerId) {
            console.log(peerId)
            const videoDom = document.getElementById(peerId);
            videoDom.remove();
        },
    },

    mounted: async function() {
        if (!this.$route.params.id) {
            alert('部屋番号が入力されていません');
            this.$router.push('/room/prepare');
        }

        try {
            // カメラ映像取得
            const stream = await navigator.mediaDevices.getUserMedia({video: true, audio: true});
            document.getElementById('my-video').srcObject = stream;
            this.localStream = stream;
        } catch(error) {
            console.log(error);
        }

        //Peer作成
        this.peer = new Peer({
            key: this.APIKey,
            debug: 3,
        });
    }
}
</script>
