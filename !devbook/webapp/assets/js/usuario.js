$("#parar-seguir").on("click", pararDeSeguir)
$("#seguir").on("click", seguir)
$("#editar-usuario").on("submit", editar)
$("#atualizar-senha").on("submit", atualizarSenha)
$("#deletar-usuario").on("click", deletarUsuario)


function pararDeSeguir() {
    const usuarioID = $(this).data("usuario-id")
    $(this).prop("disabled", true)

    $.ajax({
        url: `/usuarios/${usuarioID}/parar-de-seguir`,
        method: "POST",
    }).done(function () {
        window.location = `/usuarios/${usuarioID}`
    }).fail(function () {
        Swal.fire("Ops...", "Erro ao parar de seguir", "error")
        $("#parar-seguir").prop("disabled", false)
    })
}

function seguir() {
    const usuarioID = $(this).data("usuario-id")
    $(this).prop("disabled", true)

    $.ajax({
        url: `/usuarios/${usuarioID}/seguir`,
        method: "POST",
    }).done(function () {
        window.location = `/usuarios/${usuarioID}`
    }).fail(function () {
        Swal.fire("Ops...", "Erro ao seguir", "error")
        $("#seguir").prop("disabled", false)
    })
}

function editar(evento) {
    evento.preventDefault()

    $.ajax({
        url: "/editar-usuario",
        method: "PUT",
        data: {
            nome: $("#nome").val(),
            email: $("#email").val(),
            nick: $("#nick").val(),
        }
    }).done(function () {
        Swal.fire("Sucesso!", "Usuario atualizado com sucesso", "success").
        then(function () {
            window.location = "/perfil"
        })

    }).fail(function () {
        Swal.fire("Ops...", "Erro ao atualizar perfil", "error")
    })
}

function atualizarSenha(evento) {
    evento.preventDefault()

    const senha = $("#nova-senha").val()

    if (senha != $("#confirmar-senha").val()) {
        Swal.fire("Ops...", "As senhas não coincidem!", "warning").
        return
    }

    $.ajax({
        url: "/atualizar-senha",
        method: "PUT",
        data: {
            atual: $("#senha-atual").val(),
            nova: senha,
        }
    }).done(function () {
        Swal.fire("Sucesso!", "Senha atualizada com sucesso", "success").
        then(function () {
            window.location = "/perfil"
        })

    }).fail(function () {
        Swal.fire("Ops...", "Erro ao atualizar senha", "error")
    })
}

function deletarUsuario() {
    Swal.fire({
        title: "Atenção",
        text: "Tem certeza que deseja apar sua conta? Essa ação é irreversível",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
    }).then(function (confirmacao) {
        if (confirmacao.value) {
            $.ajax({
                url: "/deletar-usuario",
                method: "DELETE"
            }).done(function () {
                Swal.fire("Sucesso!", "Seu usuário foi excluido com sucesso", "success")
                    .then(function () {
                        window.location = "/logout"
                    })
            }).fail(function () {
                Swal.fire("Ops...", "Ocorreu um erro ao excluir seu usuário", "error")
            })
        }
    })
}