<template>
    <div>
        <section id="video-wrap"> 
            <div>
                <video id="my-video" width="400px" autoplay muted playsinline></video>
            </div>
        </section>
        <h3>peeer id: <span id="my-id"></span></h3>
        <div>
            <v-text-field
            class="d-inline-flex pa-2"
            id="room-name"
            label="Room Name"
            :rules="rules"
            hide-details="auto"
            style="width: 50%;"
            ></v-text-field>
            <v-btn id="make-call" color="primary" elevation="4">Create Room</v-btn>
            <v-btn id="end-call" color="primary" elevation="4">切断</v-btn>
        </div>
    </div>
</template>

<script>
import Peer from 'skyway-js'

export default {
    layout: "test",
    data() {
        return {
            APIKey: '5152bad7-4798-40b1-986a-a7e8f164a8a3',
            selectedAudio: '',
            selectedVideo: '',
            audios: [],
            videos: [],
            localStream: null,
            peerId: '',
            calltoid: ''
        }
    },

    methods: {
        setEventListener: function(mediaConnection) {
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
            peer.on('error', err => {
                alert(err.message);
            });
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
        }
    },
    

    mounted: async function() {
        try {
            // カメラ映像取得
            const stream = await navigator.mediaDevices.getUserMedia({video: true, audio: true});
            document.getElementById('my-video').srcObject = stream;
            this.localStream = stream;
        } catch(error) {
            console.log(error);
        }

        //Peer作成
        const peer = new Peer({
            key:   this.APIKey,
            debug: 3,
        });

        //PeerID取得
        peer.on('open', () => {
            document.getElementById('my-id').textContent = peer.id;
        });

        //グループ参加
        document.getElementById('make-call').onclick = () => {
            const roomName = document.getElementById('room-name').value;
            const mediaConnection = peer.joinRoom(roomName, {mode: 'sfu', stream: this.localStream});
            this.setEventListener(mediaConnection);
        };

        //退出
        document.getElementById('end-call').onclick = () => {
            peer.destroy();
            alert('退出しました');
        };

        //着信処理(相手から接続要求が来たタイミングで発生)
        peer.on('call', mediaConnection => {
            mediaConnection.answer(this.localStream);
            this.setEventListener(mediaConnection);
        });
    }
};
</script>
