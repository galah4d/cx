#version 130

uniform mat4 u_m44World;
uniform mat4 u_m44View;
uniform mat4 u_m44Projection;

in vec3 i_v3Position;
in vec4 i_v4Albedo;
in vec2 i_v2Texcoord;

out vec4 v_v4Albedo;
out vec2 v_v2Texcoord;

void main()
{
	gl_Position = u_m44Projection * u_m44View * u_m44World * vec4(i_v3Position, 1);

	v_v4Albedo = i_v4Albedo;
	v_v2Texcoord = i_v2Texcoord;
}
