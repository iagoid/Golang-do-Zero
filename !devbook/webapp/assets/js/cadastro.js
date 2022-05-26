$('#formulario-cadastro').on('submit', criarUsuario)

function criarUsuario(evento) {
    evento.preventDefault()

    var senha = $('#senha').val()
    var confirmarSenha = $('#confirmar-senha').val()


    if (senha != confirmarSenha) {
        Swal.fire("Ops...", "A senha não coincide com o usuário cadastrado", "error")
        return
    }

    $.ajax({
        url: "/usuarios",
        method: "POST",
        data: {
            nome: $("#nome").val(),
            email: $("#email").val(),
            nick: $("#nick").val(),
            senha: $("#senha").val(),
        }
    }).done(function () {
        $.ajax({
            url: "/login",
            method: "POST",
            data: {
                email: $("#email").val(),
                senha: $("#senha").val()
            }
        }).done(function () {
            window.location = "/home"
        }).fail(function () {
            Swal.fire("Ops...", "Seu usuário foi cadastrado, porém houve um erro ao fazer login", "error")
        })
    }).fail(function (erro) {
        Swal.fire("Ops...", "Não foi possível cadastrar este usuário", "error")
    })
}