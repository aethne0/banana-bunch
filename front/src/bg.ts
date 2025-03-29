const bg_vrt_src = `
            attribute vec2 position;
            varying vec2 fragCoord;
            void main() {
                fragCoord = position * 0.5 + 0.5; 
                gl_Position = vec4(position, 0.0, 1.0);
            }
        `;

const bg_frag_src = `
        precision mediump float;
        uniform float time;
        uniform vec2 res;
        varying vec2 fragCoord;

        vec3 blue_col = vec3(0.0,.502,.941);
        vec3 red_col = vec3(.804,.1294,.31);

        void main(){
            vec2 uv = vec2(
                (fragCoord.x - 0.5) * (res.x / res.y) + 0.5,
                fragCoord.y
            );

            float t = time * 0.001;

            // actual
            vec3 col = 0.5 + 0.5*cos(t+vec3(uv.x, uv.y, uv.x)+vec3(0,2,4));

            gl_FragColor = vec4(col, 1.0);
        }

    `;

const createShader = (gl, type, source) => {
    const shader = gl.createShader(type);
    gl.shaderSource(shader, source);
    gl.compileShader(shader);
    if (!gl.getShaderParameter(shader, gl.COMPILE_STATUS)) {
        console.error(gl.getShaderInfoLog(shader));
        gl.deleteShader(shader);
        return null;
    }
    return shader;
};

const createProgram = (gl, vertexShader, fragmentShader) => {
    const program = gl.createProgram();
    gl.attachShader(program, vertexShader);
    gl.attachShader(program, fragmentShader);
    gl.linkProgram(program);
    if (!gl.getProgramParameter(program, gl.LINK_STATUS)) {
        console.error(gl.getProgramInfoLog(program));
        gl.deleteProgram(program);
        return null;
    }
    return program;
};

const main = (canvas_id) => {
    const canvas = document.getElementById(canvas_id);
    const gl = canvas.getContext("webgl");
    if (!gl) {
        console.error("WebGL not supported");
        return;
    }

    const vertex_shader = createShader(gl, gl.VERTEX_SHADER, bg_vrt_src);
    const fragment_shader = createShader(gl, gl.FRAGMENT_SHADER, bg_frag_src);
    const program = createProgram(gl, vertex_shader, fragment_shader);

    const position_buffer = gl.createBuffer();
    gl.bindBuffer(gl.ARRAY_BUFFER, position_buffer);
    const positions = [-1, -1, 1, -1, -1, 1, -1, 1, 1, -1, 1, 1];
    gl.bufferData(gl.ARRAY_BUFFER, new Float32Array(positions), gl.STATIC_DRAW);

    const position_loc = gl.getAttribLocation(program, "position");
    gl.enableVertexAttribArray(position_loc);
    gl.vertexAttribPointer(position_loc, 2, gl.FLOAT, false, 0, 0);

    const time_loc = gl.getUniformLocation(program, "time");
    const res_loc = gl.getUniformLocation(program, "res");

    gl.viewport(
        0,
        0,
        (gl.canvas.width = window.innerWidth),
        (gl.canvas.height = window.innerHeight)
    );

    const resolution = new Float32Array(2);

    gl.useProgram(program);

    const render = (time) => {
        resolution[0] = window.innerWidth;
        resolution[1] = window.innerHeight;
        gl.uniform1f(time_loc, time);
        gl.uniform2fv(res_loc, resolution);

        gl.clear(gl.COLOR_BUFFER_BIT);
        gl.drawArrays(gl.TRIANGLES, 0, 6);
        requestAnimationFrame(render);
    };

    requestAnimationFrame(render);
};

export const initializeBg = (canvas_id) =>
    window.onload = () =>
        main(canvas_id);
