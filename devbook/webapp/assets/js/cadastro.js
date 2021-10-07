$('#formulario-cadastro').on('submit', criarUsuario)

function criarUsuario(evento) {
    evento.preventDefault()

    var senha = $('#senha').val()
    var confirmarSenha = $('#confirmar-senha').val()


    if (senha != confirmarSenha) {
        alert("Senhas não coincidem")
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
        alert("Cadastrado")
    }).fail(function (erro) {
        alert(erro)
    })
}